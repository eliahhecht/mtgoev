package main

import (
	"fmt"
	"github.com/wayn3h0/go-decimal"
	"math"
)

type PrizeSupplier interface {
	GetPrizesForPlace(place int) *decimal.Decimal
}

type PrizePayout struct {
	prizesByPlace map[int][]Pack
}

func (payout *PrizePayout) RegisterPrize(place int, prize []Pack) {
	payout.prizesByPlace[place] = prize
}

func (payout *PrizePayout) GetPrizesForPlace(place int) *decimal.Decimal {
	prize := payout.prizesByPlace[place]
	prizeValue := decimal.New(0)
	if prize != nil {
		for _, prizePack := range prize {
			prizeValue = prizeValue.Add(prizePack.ValueToSell)
		}
	}

	return prizeValue
}

type Pack struct {
	Set         PackEV
	CostToBuy   *decimal.Decimal
	ValueToSell *decimal.Decimal
}

func (pack *Pack) PackEV() float64 {
	return pack.Set.PackEV()
}

type Event struct {
	EntryPacks  []Pack
	Prizes      PrizeSupplier
	EntryFeeTix int

	// todo: assuming all events are 8-player Swiss for now
}

func (event *Event) EV(matchWinPct float64) float64 {
	valueOfOpenedCards := 0.0
	entryCost := decimal.New(float64(event.EntryFeeTix))
	for _, pack := range event.EntryPacks {
		valueOfOpenedCards += pack.PackEV()
		entryCost = entryCost.Add(pack.CostToBuy)
	}

	expectedPrizes := 0.0
	for i := 1; i <= 8; i++ {
		expectedPrizes += chanceOfPlace(i, matchWinPct) *
			event.Prizes.GetPrizesForPlace(i).Float()
	}

	fmt.Printf("voc: %f, prizes: %f, entrycost: %f\n",
		valueOfOpenedCards, expectedPrizes, entryCost.Float())

	return valueOfOpenedCards + expectedPrizes - entryCost.Float()
}

func chanceOfPlace(place int, matchWinPct float64) float64 {
	matchLossPct := 1 - matchWinPct

	switch place {
	case 1:
		return math.Pow(matchWinPct, 3) // 3-0
	case 2, 3, 4:
		return math.Pow(matchWinPct, 2) * matchLossPct // 2-1
	case 5, 6, 7:
		return math.Pow(matchLossPct, 2) * matchWinPct // 1-2
	case 8:
		return math.Pow(matchLossPct, 3) // 0-3
	default:
		panic("Unexpected place requested")
	}
}
