package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/govalues/decimal"
)

type CartItem struct {
	UserID   uuid.UUID `gorm:"column:user_id;type:uuid;primaryKey"`
	ItemID   int       `gorm:"column:item_id;type:int;primaryKey"`
	Quantity int       `gorm:"column:quantity;type:int;default:1"`
	AddedAt  time.Time `gorm:"column:added_at;type:timestamp;default:current_timestamp;primaryKey"`
}

type ActualCartItem struct {
	Product  Product
	Quantity int
}

// Optional aggregate, not persisted; useful for responses.
type Cart struct {
	Items []CartItem
	Total decimal.Decimal
}
