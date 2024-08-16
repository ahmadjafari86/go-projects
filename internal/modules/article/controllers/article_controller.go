package controllers

import (
	ArticleService "blog/internal/modules/article/services"
	"blog/pkg/html"
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
		html.Render(c, http.StatusUnprocessableEntity, "templates/errors/html/500", gin.H{"title": "Server error", "message": "Error in parsing ID param"})
		return
	}

	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Article page", "article": article})
}
