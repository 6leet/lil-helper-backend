package router

import (
	v1 "lil-helper-backend/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func InitHelperRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	HelperRouter := Router.Group("helper")

	{
		HelperRouter.POST("regist", v1.RegistHelper)
	}
	return HelperRouter
}
