package initrouter

import (
	_ "lil-helper-backend/docs"
	"lil-helper-backend/middleware"
	"lil-helper-backend/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ApiGroup := Router.Group("backend")
	router.InitAdminRouter(ApiGroup)
	router.InitHelperRouter(ApiGroup)
	router.InitCommonRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return Router
}
