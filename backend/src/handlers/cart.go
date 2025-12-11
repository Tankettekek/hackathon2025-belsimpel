package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (dbc *DBContext) HandleGetCart(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("Can't convert id to integer: %s", err))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("Can't convert id to integer: %s", err))
	}
}

func (dbc *DBContext) HandleGetCheckout(c *gin.Context) {

}

func (dbc *DBContext) HandleAddCartItem(c *gin.Context) {

}

func (dbc *DBContext) HandleRemoveCartItem(c *gin.Context) {
}
