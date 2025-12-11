package handlers

import (
	"fmt"
	"strconv"
	"time"

	"belsimel/hackathon/v2/src/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (dbc *DBContext) HandleGetCart(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("invalid user id: %s", err))
		return
	}

	// Get all cart items for the user from the database
	var cartItems []models.CartItem
	dbc.DB.Where("user_id = ?", userID).Find(&cartItems)

	var actualItems []models.ActualCartItem
	for _, item := range cartItems {
		var product models.Product
		dbc.DB.First(&product, item.ItemID)

		actualItems = append(actualItems, models.ActualCartItem{Product: product, Quantity: item.Quantity})
	}

	c.HTML(200, "cartTemplate.html", actualItems)
}

func (dbc *DBContext) HandleGetCheckout(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("invalid user id: %s", err))
		return
	}

	// Get all cart items for the user from the database
	var cartItems []models.CartItem
	dbc.DB.Where("user_id = ?", userID).Find(&cartItems)

	var actualItems []models.ActualCartItem
	for _, item := range cartItems {
		var product models.Product
		dbc.DB.First(&product, item.ItemID)

		actualItems = append(actualItems, models.ActualCartItem{Product: product, Quantity: item.Quantity})
	}

	c.HTML(200, "checkoutTemplate.html", actualItems)
}

func (dbc *DBContext) HandleAddCartItem(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user-id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("invalid user id: %s", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert product id to integer: %s", err))
		return
	}

	quantity, err := strconv.Atoi(c.DefaultQuery("quantity", "1"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert quantity to integer: %s", err))
		return
	}

	dbc.DB.Create(&models.CartItem{
		UserID:   userID,
		ItemID:   id,
		Quantity: quantity,
		AddedAt:  time.Now(),
	})
	c.Status(204)
}

func (dbc *DBContext) HandleRemoveCartItem(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user-id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("invalid user id: %s", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert product id to integer: %s", err))
		return
	}

	dbc.DB.Where("user_id = ? AND item_id = ?", userID, id).Delete(&models.CartItem{})
	c.Status(204)
}
