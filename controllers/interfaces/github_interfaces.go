package interfaces

import (
	echo "github.com/labstack/echo/v4"
)

type GitController interface {
	GetGitJobs(c echo.Context) error
}
