package models

import (
	"github.com/govalues/decimal"
)

type Product struct {
	ID          int             `gorm:"column:id;primaryKey"`
	Name        string          `gorm:"type:text;not null"`
	Description string          `gorm:"type:text;not null"`
	Price       decimal.Decimal `gorm:"type:decimal;not null"`
	Currency    string          `gorm:"type:text;not null"`
	Stock       int             `gorm:"type:int;not null"`
	Category    string          `gorm:"type:category;not null"`
	PaymentType string          `gorm:"type:payment_type;not null"`
	Attributes  []byte          `gorm:"type:jsonb"`
}
