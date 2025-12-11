package models

import (
	"github.com/govalues/decimal"
)

type PaymentType int
type Category int

const (
	OneTimePurchase PaymentType = iota
	RecurringPurchase
)

const (
	Smartphone Category = iota
	Subscription
	Prepaid
	Accessory
	Tablet
	Wearable
)

type Product struct {
	Id          int
	Stock       int
	Category    Category
	PaymentType PaymentType
	Name        string
	Price       decimal.Decimal
	Currency    string
	Description string
}
