package controllers

import (
	ArticleService "blog/internal/modules/article/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Error in parsing ID param"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Article ID is ": id})
}
