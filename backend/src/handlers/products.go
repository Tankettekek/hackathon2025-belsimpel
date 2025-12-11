package handlers

import (
	"belsimel/hackathon/v2/src/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var filterMap = map[string]string{
	"smartphone":   "smartphone",
	"subscription": "subscription",
	"prepaid":      "prepaid",
	"accessory":    "accessory",
	"tablet":       "tablet",
	"wearable":     "wearable",
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

	filterValue, hasFilter := filterMap[c.Param("filter")]

	// Query products
	var products []models.Product
	var res *gorm.DB
	if !hasFilter {
		res = dbc.DB.Offset(offset).Limit(page_size).Find(&products)
	} else {
		res = dbc.DB.Offset(offset).Limit(page_size).Where("category = ?", filterValue).Find(&products)
	}

	if res.Error != nil {
		c.AbortWithError(500, fmt.Errorf("database error: %s", res.Error))
		return
	}

	c.HTML(200, "productsTemplate.html", products)
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

	c.HTML(200, "productTemplate.html", product)
}
