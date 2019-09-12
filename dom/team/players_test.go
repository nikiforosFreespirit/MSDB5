package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

var testPlayers Players

func init() {
	var a player.Player
	a.RegisterAs("A")
	testPlayers.Add(&a)
	var b player.Player
	b.RegisterAs("B")
	b.Hand().Add(*card.MustID(33))
	testPlayers.Add(&b)
}

func TestSuccessfulFindNoErr(t *testing.T) {
	if _, p := testPlayers.Find(player.IsCardInHand(*card.MustID(33))); p == nil {
		t.Fatal("Player not found with criteria player.IsCardInHand(33)")
	}
}

func TestSuccessfulFindWithNone(t *testing.T) {
	if testPlayers.None(player.IsCardInHand(*card.MustID(33))) {
		t.Fatal("Player not found with criteria player.IsCardInHand(33)")
	}
}

func TestSuccessfulFindDataCorresponds(t *testing.T) {
	isPlayerACheck := func(p *player.Player) bool { return p.Name() == "A" }
	if _, player := testPlayers.Find(isPlayerACheck); !isPlayerACheck(player) {
		t.Fatalf("%s and %v are expected to be the same player", "A", player)
	}
}

func TestUnsuccessfulFindWithNone(t *testing.T) {
	if !testPlayers.None(player.IsCardInHand(*card.MustID(24))) {
		t.Fatal("Player should not be found")
	}
}

func TestUnsuccessfulFindWithAll(t *testing.T) {
	if testPlayers.All(player.IsCardInHand(*card.MustID(24))) {
		t.Fatal("Player should not be found")
	}
}

func TestUnsuccessfulFind(t *testing.T) {
	if _, p := testPlayers.Find(player.IsCardInHand(*card.MustID(8))); p != nil {
		t.Fatal("Player should not be found")
	}
}
