package router

import (
	v1 "lil-helper-backend/controller/api/v1"
	"lil-helper-backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitHelperRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	HelperRouter := Router.Group("helper")
	HelperRouter.Use(middleware.Jwt())
	{
		HelperRouter.GET("mission", v1.GetMission)
		HelperRouter.GET("screenshot", v1.GetScreenshot)
		HelperRouter.GET("screenshots", v1.GetScreenshots)

		HelperRouter.POST("regist", v1.RegistHelper)
		HelperRouter.POST("screenshot", v1.CreateScreenshot)
		HelperRouter.POST("screenshots/:uid", v1.UpdateScreenshot)

		HelperRouter.DELETE("screenshots/:uid", v1.DeleteScreenshot)
	}
	return HelperRouter
}
