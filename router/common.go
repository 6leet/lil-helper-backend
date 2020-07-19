package router

import (
	v1 "lil-helper-backend/controller/api/v1"
	"lil-helper-backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitCommonRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	CommonRouter := Router.Group("")
	CommonRouter.Use(middleware.Jwt())
	{
		CommonRouter.GET("helpers", v1.GetTopScoreHelpers)
		CommonRouter.GET("helpers/:uid", v1.GetTopScoreHelpersLimit)
	}
	return CommonRouter
}
