package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RequestMethodHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.Method,
	})
}

func GetGinRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	r.PATCH("/patch", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PATCH",
		})
	})
	r.HEAD("/head", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HEAD",
		})
	})
	r.OPTIONS("/options", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OPTIONS",
		})
	})
	return r
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
