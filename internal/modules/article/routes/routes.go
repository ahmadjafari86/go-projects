package routes

import (
	"blog/internal/middlewares"
	ArticleCtrl "blog/internal/modules/article/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articleController := ArticleCtrl.New()

	authGroup := router.Group("/articles")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", articleController.Create)
	}

	router.GET("/articles/:id", articleController.Show)
}
