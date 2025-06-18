package main

import (
	"backend/config"
	"backend/db"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.InitDB()

	r := gin.Default()
	routes.RegisterUserRoutes(r)
	// routes.RegisterExperimentRoutes(r)
	// routes.RegisterEventRoutes(r)
	// routes.RegisterVariantRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}