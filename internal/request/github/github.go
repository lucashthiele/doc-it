package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lucashthiele/doc-it/internal/config"
	"github.com/lucashthiele/doc-it/internal/models"
	"github.com/lucashthiele/doc-it/pkg/env"
	"github.com/lucashthiele/doc-it/pkg/logger"
	"github.com/rs/zerolog"
)

type ExchangeCodeResponse struct {
	AccessToken           string
	ExpiresIn             time.Time
	RefreshToken          string
	RefreshTokenExpiresIn time.Time
}

type UserResponse struct {
	Email string `json:"email"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

var log *zerolog.Logger = logger.Get()

type IGithubApi interface {
	ExchangeOAuthCode() error
	FillUserInfo() error
}

type GithubApi struct {
	User *models.User
}

func (ga *GithubApi) ExchangeOAuthCode() error {
	reqURL := config.GithubBaseURL + "/login/oauth/access_token"

	queryParams, err := buildQueryParams(ga.User)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	resp, err := http.PostForm(reqURL, queryParams)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	userResponse, err := parseBody(string(body))
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	ga.User.Oauth.AccessToken = userResponse.AccessToken
	ga.User.Oauth.RefreshToken = userResponse.RefreshToken
	ga.User.Oauth.ExpiresIn = userResponse.ExpiresIn
	ga.User.Oauth.RefreshTokenExpiresIn = userResponse.RefreshTokenExpiresIn

	return nil
}

func (ga *GithubApi) FillUserInfo() error {
	if ga.User.Oauth.AccessToken == "" {
		errMessage := "access token is empty"
		log.Error().Msg(errMessage)
		return fmt.Errorf(errMessage)
	}

	client := &http.Client{}
	reqURL := config.GithubApiBaseUrl + "/user"

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	req.Header.Add("Authorization", "Bearer "+ga.User.Oauth.AccessToken)
	req.Header.Add("Accept", "application/vnd.github+json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	var userResponse UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	ga.User.Email = userResponse.Email
	ga.User.Login = userResponse.Login
	ga.User.Name = userResponse.Name

	return nil
}

func parseBody(body string) (ExchangeCodeResponse, error) {
	params := strings.Split(body, "&")
	fields := make(map[string]string)

	for _, p := range params {
		keyValue := strings.Split(p, "=")
		fields[keyValue[0]] = keyValue[1]
	}

	if err, found := fields["error"]; found {
		msg := fields["error_description"]
		return ExchangeCodeResponse{}, fmt.Errorf("%s: %s", err, strings.ReplaceAll(msg, "+", " "))
	}

	expiresIn, err := time.ParseDuration(fields["expires_in"] + "s")
	if err != nil {
		return ExchangeCodeResponse{}, err
	}

	refreshTokenExpiresIn, err := time.ParseDuration(fields["refresh_token_expires_in"] + "s")
	if err != nil {
		return ExchangeCodeResponse{}, err
	}

	return ExchangeCodeResponse{
		AccessToken:           fields["access_token"],
		ExpiresIn:             time.Now().Add(expiresIn),
		RefreshToken:          fields["refresh_token"],
		RefreshTokenExpiresIn: time.Now().Add(refreshTokenExpiresIn),
	}, nil
}

func buildQueryParams(u *models.User) (url.Values, error) {
	clientID, clientSecret, err := getEnvKeys()
	if err != nil {
		return nil, err
	}
	return url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"code":          {u.Oauth.Code},
	}, nil
}

func getEnvKeys() (string, string, error) {
	clientID, err := env.GitHubClientID.GetValue()
	if err != nil {
		return "", "", err
	}
	clientSecret, err := env.GitHubClientSecret.GetValue()
	if err != nil {
		return "", "", err
	}

	return clientID, clientSecret, nil
}
