package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/wayn3h0/go-decimal"
	"testing"
)

type SetWithEV1 struct{}

func (s *SetWithEV1) PackEV() float64 {
	return 1
}

type ConstPrizesOnePack struct{}

func (p *ConstPrizesOnePack) GetPrizesForPlace(place int) *decimal.Decimal {
	return decimal.New(1)
}

func GetTestPack() Pack {
	return Pack{Set: &SetWithEV1{}, CostToBuy: decimal.New(1), ValueToSell: decimal.New(1)}
}

func TestFlatPrizeDistroComputesEVCorrectly(t *testing.T) {
	testPack := GetTestPack()
	event := Event{
		EntryPacks:  []Pack{testPack, testPack, testPack},
		EntryFeeTix: 2,
		Prizes:      &ConstPrizesOnePack{}}

	ev := event.EV(0.5)

	// cost: 3 packs + 2 tix = 5
	// voc: 3 packs * 1 = 3
	// expected prizes: 1 pack
	assert.Equal(t, -1.0, ev, "The EV of the event should be computed properly")
}
