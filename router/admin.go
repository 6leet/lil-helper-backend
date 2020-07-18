package router

import (
	v1 "lil-helper-backend/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	AdminRouter := Router.Group("admin")

	{
		AdminRouter.GET("screenshots", v1.GetScreenshots)
		AdminRouter.GET("helpers", v1.GetHelpers)
		AdminRouter.GET("missions", v1.GetMissions)
		AdminRouter.GET("missions/:date", v1.GetMissionsByDate)

		AdminRouter.POST("regist", v1.RegistAdmin)
		AdminRouter.POST("login", v1.LoginAdmin)
		AdminRouter.POST("createmission", v1.CreateMission)
		AdminRouter.POST("missions/:uid", v1.UpdateMission)
		AdminRouter.POST("setscreenshotapprove", v1.SetScreenshotApprove)

		AdminRouter.DELETE("missions/:uid", v1.DeleteMission)
		AdminRouter.DELETE("helpers/:uid", v1.BanHelper)
	}
	return AdminRouter
}
