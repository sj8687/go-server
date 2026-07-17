package main

import (
	"time"
	
	"github.com/sj8687/backend/config"
	"github.com/sj8687/backend/routes"
	"github.com/gin-contrib/cors"


	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8081", // Expo Web
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
		},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	routes.AuthRoutes(router)

	router.Run(":8080")
}
