package handlers

import (
	"belsimel/hackathon/v2/src/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		c.AbortWithError(400, fmt.Errorf("can't convert offset to integer: %s", err))
	}

	page_size, err := strconv.Atoi(c.Param("page-size"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert page size to integer: %s", err))
	}

	filter := filterMap[c.Param("filter")]
	if filter == models.Default {
		c.AbortWithError(400, fmt.Errorf("filter %s doesn't exist", c.Param("filter")))
	}

	// Query products
	var products []models.Product
	var res *gorm.DB
	if filter == models.Default {
		res = dbc.DB.Offset(offset).Limit(page_size).Find(&products)
	} else {
		res = dbc.DB.Offset(offset).Limit(page_size).Where("category = ?", filter).Find(&products)
	}

	if res.Error != nil {
		c.AbortWithError(500, fmt.Errorf("database error: %s", res.Error))
		return
	}

}

func (dbc *DBContext) HandleGetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, fmt.Errorf("can't convert id to integer: %s", err))
	}

	// get from the database with an id (Primary Key)
	var product models.Product
	result := dbc.DB.First(&product, id)
	if result.Error != nil {
		c.AbortWithError(404, fmt.Errorf("product with id %d not found", id))
		return
	}

}
