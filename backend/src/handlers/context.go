package handlers

import (
	"gorm.io/gorm"
)

type status int

const (
	Connected status = iota
	Disconnected
)

type DBContext struct {
	DB *gorm.DB
}
