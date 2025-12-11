package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/govalues/decimal"
)

type CartItem struct {
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	ItemID   int       `gorm:"type:int;primaryKey"`
	Quantity int       `gorm:"type:int;default:1"`
	AddedAt  time.Time `gorm:"type:timestamp;default:current_timestamp;primaryKey"`
}

// Optional aggregate, not persisted; useful for responses.
type Cart struct {
	Items []CartItem
	Total decimal.Decimal
}
