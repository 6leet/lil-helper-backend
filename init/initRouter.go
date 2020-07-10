package initRouter

import (
	"github.com/btwear/glx-platform/middleware"
	"github.com/btwear/glx-platform/router"
	"github.com/gin-gonic/gin"
	_ "github.com/lil-helper-backend/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	Router.Use(middleware.Logger())

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ApiGroup := Router.Group("platform")
	router.InitBaseRouter(ApiGroup)
	router.InitAdminRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	router.InitProxyRouter(ApiGroup)

	MemberGroup := Router.Group("member")
	router.InitMemberRouter(MemberGroup)

	return Router
}
