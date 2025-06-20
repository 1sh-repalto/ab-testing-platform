package main

import (
	"backend/config"
	"backend/db"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	config.LoadEnv()
	db.InitDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	routes.RegisterUserRoutes(r)
	routes.RegisterExperimentRoutes(r)
	// routes.RegisterEventRoutes(r)
	routes.RegisterVariantRoutes(r)

	log.Println("Server running on :8080")
	log.Println("Running in", gin.Mode(), "mode")

	r.Run(":8080")
}
