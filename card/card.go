package card

const (
	RARITY_C = iota
	RARITY_U
	RARITY_R
	RARITY_M
)

type Card struct {
	Price  float32
	Name   string
	Rarity int
}
