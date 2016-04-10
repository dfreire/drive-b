package middleware

import (
	"github.com/labstack/echo"
	"github.com/mgutz/logxi/v1"
)

func CreateAppMiddleware(jwtKey []byte, authorizedURL string) echo.HandlerFunc {
	return func(c *echo.Context) error {

		log.Info("UserAgent", c.Request().UserAgent())
		log.Info("RemoteAddr", c.Request().RemoteAddr)
		log.Info("RemoteAddr", c.Request().Header.Get("X-Forwarded-For"))

		return nil
	}
}
