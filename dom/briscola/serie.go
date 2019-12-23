package briscola

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

type seeder interface {
	Seed() card.Seed
}

// Serie func
func Serie(briscola seeder) set.Cards {
	serie := []uint8{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	cards := make(set.Cards, len(serie))
	for i, id := range serie {
		cards[i] = *card.MustID(id + 10*uint8(briscola.Seed()))
	}
	return cards
}