package routes

import (
	"belsimel/hackathon/v2/src/handlers"

	"github.com/gin-gonic/gin"
)

func StartRouting(dbc *handlers.DBContext) {
	router := gin.Default()

	router.GET("/products", dbc.HandleGetProducts)
	router.GET("/products/:id", dbc.HandleGetProduct)

	router.GET("/cart", dbc.HandleGetCart)
	router.GET("/cart/checkout", dbc.HandleGetCheckout)
	router.POST("/cart/add/:id", dbc.HandleAddCartItem)
	router.POST("/cart/remove/:id", dbc.HandleRemoveCartItem)

	// Runs on 8080
	router.Run()
}
