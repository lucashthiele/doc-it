package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lucashthiele/doc-it/config"
	"github.com/lucashthiele/doc-it/env"
	"github.com/lucashthiele/doc-it/logger"
	"github.com/lucashthiele/doc-it/models"
	"github.com/rs/zerolog"
)

type UserResponse struct {
	AccessToken           string
	ExpiresIn             time.Time
	RefreshToken          string
	RefreshTokenExpiresIn time.Time
}

var log *zerolog.Logger = logger.Get()

func SaveGithubUser(code string) (models.User, error) {
	user := models.User{
		Oauth: models.Oauth{
			Code: code,
		}}

	fillGithubUserData(&user)

	// todo - save info to db

	return user, nil
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

func fillGithubUserData(u *models.User) error {
	clientID, clientSecret, err := getEnvKeys()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	reqURL := config.GithubBaseURL + "/login/oauth/access_token"

	res, err := http.PostForm(reqURL,
		url.Values{
			"client_id":     {clientID},
			"client_secret": {clientSecret},
			"code":          {u.Oauth.Code},
		})
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	// todo - handle status code != 200

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	userResponse, err := parseBody(string(body))
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	u.Oauth.AccessToken = userResponse.AccessToken
	u.Oauth.RefreshToken = userResponse.RefreshToken
	u.Oauth.ExpiresIn = userResponse.ExpiresIn
	u.Oauth.RefreshTokenExpiresIn = userResponse.RefreshTokenExpiresIn

	// todo - make request for user info

	return nil
}

func parseBody(body string) (UserResponse, error) {
	params := strings.Split(body, "&")
	fields := make(map[string]string)

	for _, p := range params {
		keyValue := strings.Split(p, "=")
		fields[keyValue[0]] = keyValue[1]
	}

	if err, found := fields["error"]; found {
		msg := fields["error_description"]
		return UserResponse{}, fmt.Errorf("%s: %s", err, strings.ReplaceAll(msg, "+", " "))
	}

	expiresIn, err := time.ParseDuration(fields["expires_in"] + "s")
	if err != nil {
		return UserResponse{}, err
	}

	refreshTokenExpiresIn, err := time.ParseDuration(fields["refresh_token_expires_in"] + "s")
	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		AccessToken:           fields["access_token"],
		ExpiresIn:             time.Now().Add(expiresIn),
		RefreshToken:          fields["refresh_token"],
		RefreshTokenExpiresIn: time.Now().Add(refreshTokenExpiresIn),
	}, nil
}
