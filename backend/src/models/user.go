package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string     `gorm:"type:text;not null"`
	Email        string     `gorm:"type:text;unique;not null"`
	PasswordHash string     `gorm:"type:text"`
	CreatedAt    time.Time  `gorm:"type:timestamp;default:current_timestamp"`
	CartID       *uuid.UUID `gorm:"type:uuid"`
}
