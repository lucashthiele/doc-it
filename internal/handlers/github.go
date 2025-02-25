package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/internal/services"
)

type (
	GitHub struct {
	}

	GitHubRequest struct {
		Code string `json:"code" validate:"required"`
	}
)

// Handles the request, validates it, call services, build response
func (gh *GitHub) Post(c echo.Context) error {
	githubRequest := new(GitHubRequest)
	if err := c.Bind(githubRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err) // todo - improve response
	}

	if err := c.Validate(githubRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err) // todo - improve response
	}

	user, err := services.SaveGithubUser(githubRequest.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err) // todo - improve response
	}

	return c.JSON(http.StatusOK, user)
}
