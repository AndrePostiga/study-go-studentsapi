package controllers

import "github.com/gin-gonic/gin"

func SetupPingController(rg *gin.RouterGroup) {
	rg.GET("/ping", GetPing)
}

func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
