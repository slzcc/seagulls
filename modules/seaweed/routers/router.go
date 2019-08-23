package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/seagulls/api/v1"
)

func LRouter(router *gin.Engine) {

	api_v1 := router.Group("/api/v1")
	{
		v1.API_V1_Gateway(api_v1)
	}
}