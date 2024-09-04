package controllers

import (
	"blog/internal/modules/article/requests/articles"
	ArticleService "blog/internal/modules/article/services"
	"blog/internal/modules/user/helpers"
	"blog/pkg/converters"
	"blog/pkg/errors"
	"blog/pkg/html"
	"blog/pkg/old"
	"blog/pkg/sessions"
	"fmt"
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

	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Article", "article": article})
}

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/article/html/create", gin.H{"title": "Create article"})
}

func (controller *Controller) Store(c *gin.Context) {
	var request articles.StoreRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(c)

	article, err := controller.articleService.StoreAsUser(request, user)

	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
