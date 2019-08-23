package v1

import (
	"seagulls/controllers"
	"github.com/gin-gonic/gin"
)

// API v1 Routers
func API_V1_Gateway(router *gin.RouterGroup){
	router.GET("/", controllers.GetDemo)
}