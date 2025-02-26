package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucashthiele/doc-it/pkg/logger"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	logger := logger.Get()

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Status >= 200 && v.Status <= 299 {
				return nil
			}

			logger.Error().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("error")
			return nil
		},
	})
}

func CorsMiddleware() echo.MiddlewareFunc {
	return middleware.CORS()
}
