package router

import (
	v1 "lil-helper-backend/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")

	{
		BaseRouter.GET("refresh-token", v1.RefreshToken)

		BaseRouter.POST("regist", v1.Regist)
		BaseRouter.POST("login", v1.HelperLogin)
		BaseRouter.POST("adminlogin", v1.AdminLogin)
		BaseRouter.POST("logout", v1.Logout)
	}
	return BaseRouter
}
