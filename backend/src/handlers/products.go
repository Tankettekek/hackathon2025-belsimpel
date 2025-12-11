package handlers

import (
	"belsimel/hackathon/v2/src/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var filterMap = map[string]models.Category{
	"smartphone":   models.Smartphone,
	"subscription": models.Subscription,
	"prepaid":      models.Prepaid,
	"accessory":    models.Accessory,
	"tablet":       models.Tablet,
	"wearable":     models.Wearable,
}

func (dbc *DBContext) HandleGetProducts(c *gin.Context) {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("Can't convert offset to integer: %s", err))
	}

	page_size, err := strconv.Atoi(c.Param("page-size"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("Can't convert page size to integer: %s", err))
	}

	filter := filterMap[c.Param("filter")]
	if filter == models.Unknown {
		c.AbortWithError(400, fmt.Errorf("Filter %s doesn't exist", filter))
	}

	// Query products
}

func (dbc *DBContext) HandleGetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("Can't convert id to integer: %s", err))
	}

	// get from the database with an id
}
