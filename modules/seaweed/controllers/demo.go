package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetDemo(c *gin.Context) {
	//c.String(http.StatusOK, "Hello World")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}