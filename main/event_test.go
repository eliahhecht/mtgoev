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

type EightFourPrizeDistro struct{}

func (_ EightFourPrizeDistro) GetPrizesForPlace(place int) *decimal.Decimal {
	switch place {
	case 1:
		return decimal.New(8.0)
	case 2:
		return decimal.New(4.0)
	default:
		return decimal.New(0.0)
	}
}

func TestEightFourComputesEVCorrectlyFor50PctWin(t *testing.T) {
	testPack := GetTestPack()
	event := Event{
		EntryPacks:  []Pack{testPack, testPack, testPack},
		EntryFeeTix: 2,
		Prizes:      EightFourPrizeDistro{}}

	ev := event.EV(0.5)

	// cost: 3 packs + 2 tix = 5
	// voc: 3 packs * 1 = 3
	// expected prizes: 1.5 packs
	assert.Equal(t, -0.5, ev, "The EV of the event should be computed properly")
}

func TestEightFourComputesEVCorrectlyFor100PctWin(t *testing.T) {
	testPack := GetTestPack()
	event := Event{
		EntryPacks:  []Pack{testPack, testPack, testPack},
		EntryFeeTix: 2,
		Prizes:      EightFourPrizeDistro{}}

	ev := event.EV(1)

	// cost: 3 packs + 2 tix = 5
	// voc: 3 packs * 1 = 3
	// expected prizes: 8 packs
	assert.InEpsilon(t, 6, ev, 0.001, "The EV of the event should be computed properly")
}
