package database

import (
	"gorm.io/gorm"
)

type status int

const (
	Connected status = iota
	Disconnected
)

type Context struct {
	DB *gorm.DB
}
