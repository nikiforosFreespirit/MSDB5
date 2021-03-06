package play

import (
	"log"
	"math/rand"
	"time"

	"github.com/mcaci/msdb5/v2/app/briscola/end"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
	"github.com/mcaci/msdb5/v2/pb"
)

func Run(g struct {
	Players      team.Players
	BriscolaCard briscola.Card
	EndRound     func(*struct {
		PlayedCards  briscola.PlayedCards
		BriscolaCard briscola.Card
	}) (*pb.Index, error)
}) struct {
	OnBoard briscola.PlayedCards
} {
	playedCards := briscola.NewPlayedCards(2)
	plIdx, err := g.Players.SelectIndex(func(p player.Player) bool { return p == g.Players[0] })
	if err != nil {
		log.Fatal("didn't expect to arrive at this point")
	}

	for !end.Cond(&end.Opts{Players: g.Players}) {
		rand.Seed(time.Now().Unix())
		hnd := g.Players[plIdx].Hand()
		info := Round(&RoundOpts{
			PlHand:       hnd,
			PlIdx:        plIdx,
			CardIdx:      uint8(rand.Intn(len(*hnd))),
			PlayedCards:  playedCards,
			NPlayers:     uint8(len(g.Players)),
			BriscolaCard: g.BriscolaCard,
			EndRound:     g.EndRound,
		})
		playedCards = info.OnBoard
		plIdx = info.NextPl
		if !info.NextRnd {
			continue
		}
		briscola.Collect(playedCards, (g.Players)[plIdx])
	}
	return struct{ OnBoard briscola.PlayedCards }{
		OnBoard: *playedCards,
	}
}
