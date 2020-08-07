package router

import (
	"fmt"
	"lil-helper-backend/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	prePath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(prePath)

	FileRouter := Router.Group("")
	FileRouter.Use(middleware.HelperJwt())
	{
		FileRouter.StaticFS("/files", http.Dir(prePath+"/files"))
	}
	return FileRouter
}
