package driveb

import "github.com/labstack/echo"

type appMiddleware struct{}

func GlobalMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			return next(c)
		}
	}
}
