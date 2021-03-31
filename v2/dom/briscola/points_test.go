package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

func TestPoints(t *testing.T) {
	if Points(MustID(1)) != Points(MustID(21)) {
		t.Fatal("unexpected result")
	}
}

func TestScore(t *testing.T) {
	fakeHand := *set.NewMust(1, 2, 3)
	if score1 := Score(fakeHand); score1 != 21 {
		t.Fatal("Points string should contain the total of 21")
	}
}
func TestSerie(t *testing.T) {
	serie := Serie(card.MustID(1))
	expectedN := []uint8{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	for i, p := range serie {
		if p.Number() != expectedN[i] {
			t.Errorf("Unexpected value at index %d: Actual %d, Expected %d", i, p.Number(), expectedN[i])
		}
	}
}
