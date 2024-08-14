package routes

import (
	ArticleRoutes "blog/internal/modules/article/routes"
	HomeRoutes "blog/internal/modules/home/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	HomeRoutes.Routes(router)
	ArticleRoutes.Routes(router)
}
