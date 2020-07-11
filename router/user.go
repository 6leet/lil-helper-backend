package router

import "github.com/gin-gonic/gin"

func InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	UserRouter := Router.Group("admin")

	{

	}
	return UserRouter
}
