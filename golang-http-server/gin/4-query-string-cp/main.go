package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// TODO: answer here
	router.GET("/movie", func(ctx *gin.Context) {
		genre := ctx.DefaultQuery("genre", "general")
		country := ctx.Query("country")
		result := fmt.Sprintf("Here result of query of movie with genre %v", genre)
		if country != "" {
			result = fmt.Sprintf("%v in country %v", result, country)
		}
		ctx.String(http.StatusOK, result)
	})

	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}