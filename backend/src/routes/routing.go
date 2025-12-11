package routes

import (
	"belsimel/hackathon/v2/src/handlers"

	"github.com/gin-gonic/gin"
)

func StartRouting(dbc *handlers.DBContext) {
	router := gin.Default()

	router.GET("/products", dbc.HandleGetProducts)
	router.GET("/products/:id", dbc.HandleGetProduct)

	router.GET("/cart/:id", dbc.HandleGetCart)
	router.GET("/cart/:id/checkout", dbc.HandleGetCheckout)
	router.POST("/cart/:id/add/:id", dbc.HandleAddCartItem)
	router.POST("/cart/:id/remove/:id", dbc.HandleRemoveCartItem)

	// Runs on 8080
	router.Run()
}
