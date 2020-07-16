package router

import "github.com/gin-gonic/gin"

func InitHelperRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	HelperRouter := Router.Group("helper")

	{

	}
	return HelperRouter
}
