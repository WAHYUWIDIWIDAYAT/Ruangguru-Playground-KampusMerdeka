package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/:url", urlHandler.Get)
	r.POST("/", urlHandler.Create)
	r.POST("/custom", urlHandler.CreateCustom)

	return r // TODO: replace this
}
