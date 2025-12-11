package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func HandleGetProducts(c *gin.Context) {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("Can't convert offset to integer: %s", err))
	}

	page_size, err := strconv.Atoi(c.Param("page-size"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("Can't convert page size to integer: %s", err))
	}

}

func HandleGetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("Can't convert id to integer: %s", err))
	}

	// get from the database with an id
}
