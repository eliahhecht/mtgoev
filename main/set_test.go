package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/wayn3h0/go-decimal"
	"testing"
)

func makeCard(rarity int, price float64) Card {
	return Card{Rarity: rarity, Price: decimal.New(price)}
}

func TestSetEvCalculatesCorrectly(t *testing.T) {
	set := Set{
		Cards: []Card{
			makeCard(RARITY_C, 1),
			makeCard(RARITY_C, 0),
			makeCard(RARITY_U, 2),
			makeCard(RARITY_R, 3),
			makeCard(RARITY_M, 4),
			makeCard(RARITY_BASIC, 99)}}

	ev := set.PackEV()

	assert.Equal(t, 14.625, ev, "The EV of a pack should be calculated correctly")
}
