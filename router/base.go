package router

import (
	v1 "lil-helper-backend/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")

	{
		BaseRouter.GET("refresh-token", v1.RefreshToken)

		BaseRouter.POST("login", v1.Login)
	}
	return BaseRouter
}
