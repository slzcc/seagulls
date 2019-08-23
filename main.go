package main

import (
	LRouter "seagulls/routers"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	LRouter.LRouter(router)
	router.Run(":8000")
}