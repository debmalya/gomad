package router

import (
	"github.com/debmalya/gomad/controllers"
	"github.com/labstack/echo/v4"
)

func setGitServices(e *echo.Echo) {
	gitHubController := &controllers.GitController{}
	e.GET(path+version+"/:skill/:pageNo", gitHubController.GetGitHubJobs)
	// log.Info(squadGetLogInfo, squadUri+"/getsquadpendingdetails")
}
