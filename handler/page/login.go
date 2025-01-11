package page

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/env"
	"github.com/lucashthiele/doc-it/view/page/login"
)

type LoginHandler struct{}

func (lh LoginHandler) Show(c echo.Context) error {
	gitHubClientID := env.GitHubClientID.GetValue()

	return Render(c, login.Show(gitHubClientID))
}
