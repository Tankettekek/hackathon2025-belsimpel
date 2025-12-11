package routes

import (
	"belsimel/hackathon/v2/src/handlers"
	"github.com/gin-gonic/gin"
)

func StartRouting() {
	router := gin.Default()

	router.GET("/products", handlers.HandleGetProducts)
	router.GET("/products/:id", handlers.HandleGetProduct)

	router.GET("/cart", handlers.HandleGetCart)
	router.GET("/cart/checkout", handlers.HandleGetCheckout)

	router.POST("/cart/add/:id", handlers.HandleAddCartItem)
	router.POST("/cart/remove/:id", handlers.HandleRemoveCartItem)

	// Runs on 8080
	router.Run()
}
