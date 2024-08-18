package controllers

import (
	"blog/internal/modules/user/requests/auth"
	userServices "blog/internal/modules/user/services"
	"blog/pkg/html"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService userServices.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: userServices.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := controller.userService.Create(request)

	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	fmt.Printf("User created successfully with a name %s", user.Name)
	c.Redirect(http.StatusFound, "/")
}
