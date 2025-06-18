package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"backend/config"
	"backend/db"
	"backend/routes"
)

func main() {
	config.LoadEnv()
	db.InitDB()

	r := gin.Default()
	routes.RegisterExperimentRoutes(r)
	routes.RegisterEventRoutes(r)
	routes.RegisterVariantRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}