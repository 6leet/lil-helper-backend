package router

import (
	v1 "lil-helper-backend/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	AdminRouter := Router.Group("admin")

	{
		AdminRouter.GET("helloadmin", v1.HelloAdmin)

		AdminRouter.POST("regist", v1.RegistUser)
	}
	return AdminRouter
}
