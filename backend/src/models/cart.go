package models

import (
	"github.com/govalues/decimal"
)

type CartItem struct {
	Product  Product
	Quantity int `gorm:"type:int;default:1"`
}

// Optional aggregate, not persisted; useful for responses.
type Cart struct {
	Items []CartItem
	Total decimal.Decimal
}
