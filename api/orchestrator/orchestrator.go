package orchestrator

import (
	"log"
	"strconv"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/auction"
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Action interface
func (g *Game) Action(request, origin string) ([]display.Info, []display.Info, error) {
	data := strings.Split(string(request), "#")
	var err error
	switch data[0] {
	case "Join":
		err = g.Join(data[1], origin)
	case "Auction":
		err = g.RaiseAuction(data[1], origin)
	case "Companion":
		err = g.Nominate(data[1], data[2], origin)
	case "Card":
		err = g.Play(data[1], data[2], origin)
	}
	pInfo, err := g.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	return g.info.Info(), pInfo.Info(), err
}

// RaiseAuction func
func (g *Game) RaiseAuction(score, origin string) error {
	p, err := g.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	if err == nil {
		prevScore := g.info.AuctionScore()
		currentScore, err := strconv.Atoi(score)
		if err != nil {
			log.Printf("Error was raised during auction: %v\n", err)
		}
		auction.Update(0, prevScore, uint8(currentScore), p.SetAuctionScore)
		auction.Update(prevScore, prevScore, uint8(currentScore), g.info.SetAuctionScore)
	}
	return err
}

// Play func
func (g *Game) Play(number, seed, origin string) error {
	p, err := g.Players().Find(func(p *player.Player) bool { return p.Host() == origin })
	if err == nil {
		c, err := p.Play(number, seed)
		if err == nil {
			roundHasEnded := g.info.PlayedCardIs(c)
			if roundHasEnded {
				playerIndex := briscola.IndexOfWinningCard(*g.info.PlayedCards(), g.companion.Card().Seed())
				g.info.PlayedCards().Move(g.Players()[playerIndex].Pile())
			}
		}
	}
	return err
}

// Nominate func
func (g *Game) Nominate(number, seed, origin string) error {
	card, err := card.Create(number, seed)
	if err == nil {
		p, err := g.Players().Find(func(p *player.Player) bool { return p.Has(card) })
		if err == nil {
			g.companion = *companion.New(card, p)
		}
	}
	return err
}

// Join func
func (g *Game) Join(name, origin string) error {
	nextPlayerJoining := func(p *player.Player) bool { return p.Name() == "" }
	p, err := g.Players().Find(nextPlayerJoining)
	if err == nil {
		p.Join(name, origin)
		if _, errNext := g.Players().Find(nextPlayerJoining); errNext != nil {
			g.statusInfo = scoreAuction
		}
	} else {
		log.Println("All players have joined, no further players are expected: " + err.Error())
		log.Println(g.Players())
	}
	return err
}
