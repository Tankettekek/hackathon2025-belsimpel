package main

import (
	"belsimel/hackathon/v2/src/handlers"
	"belsimel/hackathon/v2/src/routes"
)

func main() {
	// Application entry point

	db, err := handlers.ConnectDatabase()
	if err != nil {
		// Handle the error appropriately
		panic("Failed to connect to database: " + err.Error())
	}

	context := &handlers.DBContext{DB: db}

	routes.StartRouting(context)
}
