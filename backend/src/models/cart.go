package models

import (
	"github.com/google/uuid"
	"github.com/govalues/decimal"
	"time"
)

type CartItem struct {
	ID       uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	ItemId   int       `gorm:"column:item_id;type:int;default:1"`
	Quantity int       `gorm:"column:quantity;type:int;default:1"`
	AddedAt  time.Time `gorm:"column:item_id;type:int;default:1"`
}

type ActualCartItem struct {
	Product  Product
	Quantity int `gorm:"type:int;default:1"`
}

// Optional aggregate, not persisted; useful for responses.
type Cart struct {
	Items []ActualCartItem
	Total decimal.Decimal
}
