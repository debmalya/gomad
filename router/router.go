package router

import (
	"github.com/labstack/echo/v4"
)

const (
	path    = "/git"
	version = "/v0"
)

// InitRoutes to initialize routes
func InitRoutes(e *echo.Echo) {
	setGitServices(e)
}
