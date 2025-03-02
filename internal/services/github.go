package services

import (
	"github.com/lucashthiele/doc-it/internal/models"
	"github.com/lucashthiele/doc-it/internal/request/github"
	"github.com/lucashthiele/doc-it/pkg/logger"
	"github.com/rs/zerolog"
)

var log *zerolog.Logger = logger.Get()

func SaveGithubUser(code string) (models.User, error) {
	user := models.User{
		Oauth: models.Oauth{
			Code: code,
		},
	}
	var err error
	githubReq := &(github.GithubApi{User: &user})

	err = githubReq.ExchangeOAuthCode()
	if err != nil {
		return models.User{}, err
	}

	err = githubReq.FillUserInfo()
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
