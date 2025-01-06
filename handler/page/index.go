package page

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/view/page/index"
)

type IndexHandler struct {
}

func (ih IndexHandler) Show(c echo.Context) error {
	return Render(c, index.Show("Lucas"))
}
