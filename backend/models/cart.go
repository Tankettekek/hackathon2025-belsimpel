package models

import (
	"github.com/govalues/decimal"
)

type Cart struct {
	Products []Product
	Total    decimal.Decimal
}
