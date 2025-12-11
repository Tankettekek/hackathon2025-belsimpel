package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id            uuid.UUID
	name          string
	email         string
	password_hash string
	created_at    time.Time
}
