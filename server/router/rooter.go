package router

import "github.com/gin-gonic/gin"

// Routers returns *gin.Engine with Router
func Routers() *gin.Engine {
	var Router = gin.Default()
	return Router
}
