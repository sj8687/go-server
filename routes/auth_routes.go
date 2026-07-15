package routes

import (
	"github.com/sj8687/backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	auth := router.Group("/auth")

	auth.POST("/login", controllers.Login)
}