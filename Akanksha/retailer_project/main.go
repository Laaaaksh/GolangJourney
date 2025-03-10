package main

import (
	"retailer_project/config"
	"retailer_project/routes"
	"log"
)

func main() {
	config.ConnectDatabase()
	config.ConnectRedis() // Initialize Redis
	r := routes.SetupRouter()
	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
