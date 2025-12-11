package main

import (
	"belsimel/hackathon/v2/src/database"
	"belsimel/hackathon/v2/src/routes"
)

func main() {
	// Application entry point

	db, err := database.ConnectDatabase()
	if err != nil {
		// Handle the error appropriately
		panic("Failed to connect to database: " + err.Error())
	}

	context := &database.Context{DB: db}

	routes.StartRouting(context)
}
