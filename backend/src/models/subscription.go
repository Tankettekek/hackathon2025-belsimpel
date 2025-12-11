package models

import (
	"time"

	"github.com/google/uuid"
)

// Aligns with active_subscriptions table
type ActiveSubscription struct {
	UserID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	ProductID     int        `gorm:"type:int;primaryKey"`
	StartDate     time.Time  `gorm:"type:timestamp;default:current_timestamp"`
	EndDate       *time.Time `gorm:"type:timestamp"`
	AdjustedPrice float64    `gorm:"type:decimal"`
}
