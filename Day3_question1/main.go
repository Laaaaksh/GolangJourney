package main

import (
	"apicall/config"
	"apicall/routes"
)

func main() {
	config.ConnectDatabase()
	config.MigrateDatabase()

	r := routes.SetupRouter()
	r.Run(":8080") // Start the server on port 8080
}
