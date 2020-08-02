package router

import (
	"lil-helper-backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	FileRouter := Router.Group("")
	FileRouter.Use(middleware.HelperJwt())
	{
		FileRouter.StaticFS("/files", http.Dir("/Users/leolee/Documents/job/lil-helper-backend/files"))
	}
	return FileRouter
}
