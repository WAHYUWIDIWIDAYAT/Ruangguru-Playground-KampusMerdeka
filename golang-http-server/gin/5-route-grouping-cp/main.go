package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": movies,
	})
}

var MovieGetHandler = func(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, ok := movies[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	movie := router.Group("/movie")
	{
		movie.GET("/list", MovieListHandler)
		movie.GET("/get/:id", MovieGetHandler)
	}
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}