package main

import (
	"github.com/gin-gonic/gin"
)

// Dari contoh yang telah diberikan, buatlah HTTP server sederhana dengan menggunakan Gin.
// Buatlah server yang memiliki address localhost dengan port 3000.
// Buatlah route "/hello" yang menampilkan JSON {"message": "hello"} dan "/world" yang menampilkan JSON {"message": "world"}

func GetGinRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})

	})
	r.GET("/world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "world",
		})
	})
	return r
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
