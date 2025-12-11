package models

import (
	"github.com/govalues/decimal"
	"gorm.io/gorm"
)

type Cart struct {
	Products []Product
	Total    decimal.Decimal
}
