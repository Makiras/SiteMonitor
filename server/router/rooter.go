package router

import {
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
}

// Routers returns *gin.Engine with Router
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	PublicGroup := Router.Group("")
	PrivateGroup := Router.Group("")
	return Router
}
