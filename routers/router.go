package routers

import (
	"seagulls/api/v1"
	"github.com/gin-gonic/gin"
)

func LRouter(router *gin.Engine) {

	api_v1 := router.Group("/api/v1")
	{
		v1.API_V1_Gateway(api_v1)
	}
}