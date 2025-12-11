package models

import (
	"time"

	"github.com/google/uuid"
)

// Aligns with orders table
type Order struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID int       `gorm:"type:int;primaryKey"`
	OrderDate time.Time `gorm:"type:timestamp;default:current_timestamp;primaryKey"`
	IsActive  bool      `gorm:"type:boolean;default:false"`
}
