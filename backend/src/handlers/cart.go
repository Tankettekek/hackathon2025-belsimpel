package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (dbc *DBContext) HandleGetCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert id to integer: %s", err))
	}

	_ = id
	c.JSON(200, gin.H{"items": []any{}, "total": 0})
}

func (dbc *DBContext) HandleGetCheckout(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert id to integer: %s", err))
	}

	_ = id
	c.JSON(200, gin.H{"ok": true})
}

func (dbc *DBContext) HandleAddCartItem(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert id to integer: %s", err))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert id to integer: %s", err))
	}

	_ = user_id
	_ = id
	c.Status(204)
}

func (dbc *DBContext) HandleRemoveCartItem(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert id to integer: %s", err))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert id to integer: %s", err))
	}

	_ = user_id
	_ = id
	c.Status(204)
}
