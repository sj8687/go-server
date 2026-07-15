package main

import (
	"github.com/sj8687/backend/config"
	"github.com/sj8687/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	routes.AuthRoutes(router)

	router.Run(":8080")
}
