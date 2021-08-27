package utils

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Authorization() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(s1, s2 string, c echo.Context) (bool, error) {
		if s1 == Getenv("username") && s2 == Getenv("password") {
			return true, nil
		} else {
			return false, nil
		}
	})
}
