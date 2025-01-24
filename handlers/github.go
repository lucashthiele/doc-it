package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/logger"
)

type (
	GitHub struct {
	}

	GitHubRequest struct {
		Code string `json:"code" validate:"required"`
	}
)

func (gh *GitHub) Post(c echo.Context) error {
	logger := logger.Get()

	githubRequest := new(GitHubRequest)
	if err := c.Bind(githubRequest); err != nil {
		return err
	}

	if err := c.Validate(githubRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Code not provided")
	}

	logger.Info().Msg("Received code: " + githubRequest.Code)

	return c.JSON(http.StatusOK, "User saved")
}
