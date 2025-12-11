package routes

import (
	"github.com/gin-gonic/gin"
)

func startRouting() {
	router := gin.Default()

	// Runs on 8080
	router.Run()
}
