package models

import (
	"github.com/govalues/decimal"
)

type Product struct {
	ID          int             `gorm:"column:id;primaryKey"`
	Name        string          `gorm:"column:name;type:text;not null"`
	Description string          `gorm:"column:description;type:text;not null"`
	Price       decimal.Decimal `gorm:"column:price;type:decimal;not null"`
	Currency    string          `gorm:"column:currency;type:text;not null"`
	Stock       int             `gorm:"column:stock;type:int;not null"`
	Category    string          `gorm:"column:category;type:category;not null"`
	PaymentType string          `gorm:"column:payment_type;type:payment_type;not null"`
	Attributes  []byte          `gorm:"column:attributes;type:jsonb"`
}
