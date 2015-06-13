package main

import "github.com/wayn3h0/go-decimal"

const (
	RARITY_C = iota
	RARITY_U
	RARITY_R
	RARITY_M
	RARITY_BASIC
)

type Card struct {
	Price  *decimal.Decimal
	Name   string
	Rarity int
}
