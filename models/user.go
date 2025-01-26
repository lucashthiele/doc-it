package models

import "time"

type (
	Oauth struct {
		Code                  string
		AccessToken           string
		ExpiresIn             time.Time
		RefreshToken          string
		RefreshTokenExpiresIn time.Time
	}

	User struct {
		Id       string
		Username string
		Email    string
		Oauth    Oauth
	}
)
