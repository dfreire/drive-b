package driveb

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type appMiddleware struct{}

func AppMiddleware(authorizedHosts []string) echo.MiddlewareFunc {
	hostMap := make(map[string]bool)
	for _, host := range authorizedHosts {
		hostMap[host] = true
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			host1 := c.Request().Host()
			host2 := c.Request().(*standard.Request).Request.URL.Host

			log.Println("Host from request:", host1)
			log.Println("Host from url:", host2)

			// log.Println("RemoteAddr:", c.Request().RemoteAddress())
			// log.Println("X-Forwarded-For:", c.Request().Header().Get("X-Forwarded-For"))

			return next(c)
		}
	}
}
