package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/deck"
)

type Board struct {
	deck deck.Deck
}

func New() *Board {
	var b Board
	b.deck = deck.New()
	return &b
}

func (b *Board) Deck() deck.Deck {
	return b.deck
}

func TestBoardHasADeck(t *testing.T) {
	b := New()
	if b.Deck() == nil {
		t.Fatal("The board has no Deck")
	}

}

func TestBoardsDeckReferenceIsTheSame(t *testing.T) {
	b := New()
	deck1 := b.Deck()
	deck2 := b.Deck()
	if deck1 != deck2 {
		t.Fatal("The deck is not the same each time is retrieved")
	}
}
