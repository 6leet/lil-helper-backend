package router

import (
	v1 "lil-helper-backend/controller/api/v1"
	"lil-helper-backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitHelperRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	HelperRouter := Router.Group("helper")
	HelperRouter.Use(middleware.HelperJwt())
	{
		HelperRouter.GET("mission", v1.GetMission)
		HelperRouter.GET("screenshots", v1.GetScreenshots)

		HelperRouter.POST("screenshot", v1.CreateScreenshot)

		HelperRouter.DELETE("screenshots/:uid", v1.DeleteScreenshot)
	}
	return HelperRouter
}
