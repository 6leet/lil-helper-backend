package router

import (
	v1 "lil-helper-backend/controller/api/v1"
	"lil-helper-backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	AdminRouter := Router.Group("admin")
	AdminRouter.Use(middleware.Jwt())
	{
		// to do: missions are now without route path

		AdminRouter.GET("screenshots", v1.GetAllScreenshots)
		AdminRouter.GET("helpers", v1.GetHelpers)
		AdminRouter.GET("missions", v1.GetMissions)
		AdminRouter.GET("reorganize", v1.ReorganizeMission)

		AdminRouter.POST("regist", v1.RegistAdmin)
		AdminRouter.POST("mission", v1.CreateMission)
		AdminRouter.POST("missions/:uid", v1.UpdateMission)
		AdminRouter.POST("screenshots/:uid", v1.SetScreenshotApprove)

		AdminRouter.DELETE("missions/:uid", v1.DeleteMission)
		AdminRouter.DELETE("helpers/:uid", v1.BanHelper)
	}
	return AdminRouter
}
