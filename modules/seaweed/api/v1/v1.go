package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/seagulls/seaweed/controllers"
)

// API v1 Routers
func API_V1_Gateway(router *gin.RouterGroup){
	router.GET("/", controllers.GetDemo)
}