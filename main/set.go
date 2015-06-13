package main

import "github.com/wayn3h0/go-decimal"

type Set struct {
	Cards []Card
}

func (s *Set) PackEV() *decimal.Decimal {
	avgCommonPrice := s.getAvgPrice(RARITY_C)
	avgUncommonPrice := s.getAvgPrice(RARITY_U)
	avgRarePrice := s.getAvgPrice(RARITY_R)
	avgMythicPrice := s.getAvgPrice(RARITY_M)

	return decimal.New(11).Mul(avgCommonPrice).
		Add(decimal.New(3).Mul(avgUncommonPrice)).
		Add(decimal.New(7 / 8.0).Mul(avgRarePrice)).
		Add(decimal.New(1 / 8.0).Mul(avgMythicPrice))
}

func (s *Set) getAvgPrice(rarity int) *decimal.Decimal {
	sum := decimal.New(0)
	count := 0.0
	for _, card := range s.Cards {
		if card.Rarity == rarity {
			count++
			sum = sum.Add(card.Price)
		}
	}
	return sum.Div(decimal.New(count))
}
