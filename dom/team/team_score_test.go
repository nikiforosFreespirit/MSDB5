package team

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
)

type mockPiler struct{}

func (mockPiler) Pile() *set.Cards {
	return set.NewMust(1)
}

func TestTeam1(t *testing.T) {
	fakePlayer := new(mockPiler)
	if score1, _ := Score(fakePlayer, nil, []Piler{fakePlayer}); score1 != 11 {
		t.Fatal("Points string should contain the total of 11")
	}
}

func TestTeam2(t *testing.T) {
	if _, score2 := Score(nil, nil, []Piler{new(mockPiler)}); score2 != 11 {
		t.Fatal("Points string should contain the total of 11")
	}
}
