package seaweed

import (
	"github.com/gin-gonic/gin"
	LRouter "github.com/slzcc/seagulls/seaweed/routers"
)

func main(){
	router := gin.Default()
	LRouter.LRouter(router)
	router.Run(":8000")
}