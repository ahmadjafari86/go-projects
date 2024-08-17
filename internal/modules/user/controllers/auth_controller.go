package controllers

import (
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "registration done."})
}
