package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Country string
	Age     int
}

var data = map[int]User{
	1: {
		Name:    "Roger",
		Country: "United States",
		Age:     70,
	},
	2: {
		Name:    "Tony",
		Country: "United States",
		Age:     40,
	},
	3: {
		Name:    "Asri",
		Country: "Indonesia",
		Age:     30,
	},
}

func ProfileHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)
		if result, ok := data[id]; ok {
			c.String(http.StatusOK, "Name: %v, Country: %v, Age: %v",
				result.Name, result.Country, result.Age)
			return
		}
		c.String(http.StatusNotFound, "data not found")
	}
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/profile/:id", ProfileHandler())
	return r
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}