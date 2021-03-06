package ml

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type mlInformer interface {
	AuctionScore() *auction.Score
	team.Callers
	CurrentPlayer() *player.Player
	LastPlayer() *player.Player
	PlayedCard() card.Item
	Phase() phase.ID
}

// Write func
func Write(g mlInformer) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		io.WriteString(os.Stdout, err.Error())
	}
	defer f.Close()
	switch g.Phase() {
	case phase.PlayingCards:
		if g.PlayedCard().Number() != 0 {
			io.WriteString(f, fmt.Sprintf("%s:%d\n", g.LastPlayer().Name(), g.PlayedCard().ToID()))
		}
	case phase.End:
		io.WriteString(f, fmt.Sprintf("%s:%d\n", g.LastPlayer().Name(), g.PlayedCard().ToID()))
		io.WriteString(f, fmt.Sprintf("%s\n", g.CurrentPlayer().Name()))
		io.WriteString(f, fmt.Sprintf("%s\n", g.Caller().Name()))
		io.WriteString(f, fmt.Sprintf("%s\n", g.Companion().Name()))
		io.WriteString(f, fmt.Sprintf("%d\n", *(g.AuctionScore())))
	}
}
