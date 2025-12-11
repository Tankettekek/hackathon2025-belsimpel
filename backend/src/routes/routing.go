package routes

import (
	"belsimel/hackathon/v2/src/handlers"

	"github.com/gin-gonic/gin"
)

func StartRouting(dbc *handlers.DBContext) {
	router := gin.Default()

	router.LoadHTMLGlob("../frontend/templates")

	router.GET("/products", dbc.HandleGetProducts)
	router.GET("/products/:id", dbc.HandleGetProduct)

	router.GET("/cart/:id", dbc.HandleGetCart)
	router.GET("/cart/:id/checkout", dbc.HandleGetCheckout)
	router.POST("/cart/:user-id/add/:id", dbc.HandleAddCartItem)
	router.POST("/cart/:user-id/remove/:id", dbc.HandleRemoveCartItem)

	// Runs on 8080
	router.Run()
}
