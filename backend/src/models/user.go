package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string     `gorm:"column:name;type:text;not null"`
	Email        string     `gorm:"column:email;type:text;unique;not null"`
	PasswordHash string     `gorm:"column:password_hash;type:text"`
	CreatedAt    time.Time  `gorm:"column:created_at;type:timestamp;default:current_timestamp"`
	CartID       *uuid.UUID `gorm:"column:cart_id;type:uuid"`
}
