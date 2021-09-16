package utils

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "uri:${uri},status:${status}\n",
	})
}
