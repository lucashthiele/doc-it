package page

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/view/page/login"
)

type LoginCallbackHandler struct{}

func (lc LoginCallbackHandler) Get(c echo.Context) error {
	authCode := c.QueryParams().Get("code")

	return Render(c, login.Callback(authCode))
}
