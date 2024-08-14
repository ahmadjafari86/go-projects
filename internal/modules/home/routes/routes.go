package routes

import (
	HomeCtrl "blog/internal/modules/home/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	homeController := HomeCtrl.New()
	router.GET("/", homeController.Index)
}
