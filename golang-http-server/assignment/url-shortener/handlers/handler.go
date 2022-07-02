package handlers

import (
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	u := c.Param("url")
	if u == "" {
		c.JSON(400, gin.H{
			"error": "url required",
		})
		return
	}
	urlEntity, err := h.repo.Get(u)
	if err != nil {
		if err == entity.ErrURLNotFound {
			c.String(404, "url not found")
		}
		c.Redirect(302, urlEntity.LongURL)
		return
	}
	c.JSON(200, gin.H{"Data": urlEntity})
}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	var reqBody entity.URL

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.String(400, "Bad request")
		return
	}
	urlEntity, err := h.repo.Create(reqBody.LongURL)
	if err != nil {
		c.String(500, "Internal server error")
		return
	}
	c.JSON(200, gin.H{"Data": urlEntity})

}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	var reqBody entity.URL

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.String(400, "Bad request")
		return
	}

	urlEntity, err := h.repo.CreateCustom(reqBody.LongURL, reqBody.ShortURL)
	if err != nil {
		c.String(500, "Internal server error")
		return
	}
	c.JSON(200, gin.H{"Data": urlEntity})
}
