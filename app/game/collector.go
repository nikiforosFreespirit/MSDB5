package game

import (
	"container/list"
	"io"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type collector interface {
	Caller() *player.Player
	Briscola() card.Item
	Companion() *player.Player
	Players() team.Players
	LastPlaying() *list.List
	Lang() language.Tag
	IsSideUsed() bool
	SideDeck() *set.Cards
}

func collect(g collector) {
	highbriscolaCard := briscola.Serie(g.Briscola())
	for _, card := range highbriscolaCard {
		_, p := g.Players().Find(player.IsCardInHand(card))
		if p == nil { // no one has card
			continue
		}
		for _, pl := range g.Players() {
			move(pl.Hand(), p.Pile())
		}
		if g.IsSideUsed() {
			move(g.SideDeck(), p.Pile())
		}
		track.Player(g.LastPlaying(), p)
		printer := message.NewPrinter(g.Lang())
		team := printer.Sprintf("Callers")
		if p != g.Caller() && p != g.Companion() {
			team = printer.Sprintf("Others")
		}
		for _, pl := range g.Players() {
			io.WriteString(pl, printer.Sprintf("The end - %s team has all briscola cards", team))
		}
		break
	}
}
