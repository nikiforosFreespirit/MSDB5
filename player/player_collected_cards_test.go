package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card/set"
)

func TestPlayerHasNoCardsAtStartGame(t *testing.T) {
	if p := New(); len(*p.Pile()) > 0 {
		t.Fatal("Player should not have cards at creation")
	}
}

func TestPlayerWinsCards(t *testing.T) {
	p := New()
	cards := set.Cards{1, 2, 3, 4, 5}
	p.Collect(cards)

	if collectedCards := p.Pile(); collectedCards == nil {
		t.Fatalf("Player should have %v but has %v", cards, collectedCards)
	}
}