package routes

import (
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
			"title":   "Home page",
			"appName": viper.Get("App.Name"),
		})
	})
	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title":   "About page",
			"appName": viper.Get("App.Name"),
		})
	})
}
