package main

type Set struct {
	Cards []Card
}

type PackEV interface {
	PackEV() float64
}

func (s *Set) PackEV() float64 {
	avgCommonPrice := s.getAvgPrice(RARITY_C)
	avgUncommonPrice := s.getAvgPrice(RARITY_U)
	avgRarePrice := s.getAvgPrice(RARITY_R)
	avgMythicPrice := s.getAvgPrice(RARITY_M)

	return 11*avgCommonPrice +
		3*avgUncommonPrice +
		(7/8.0)*avgRarePrice +
		(1/8.0)*avgMythicPrice
}

func (s *Set) getAvgPrice(rarity int) float64 {
	sum := 0.0
	count := 0.0
	for _, card := range s.Cards {
		if card.Rarity == rarity {
			count++
			sum = sum + card.Price.Float()
		}
	}
	return sum / count
}
