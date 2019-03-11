package orchestrator

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) play(request, origin string) (all []display.Info, me []display.Info, err error) {
	_, err = cardAction(request)
	if err != nil {
		return
	}
	playerInTurn := g.playerInTurn
	roundMayEnd := len(*g.info.PlayedCards()) >= 4
	if roundMayEnd {
		info := g.playEndRoundData(request, origin)
		err = g.playPhase(info)
	} else {
		info := g.playData(request, origin)
		err = g.playPhase(info)
	}
	if g.phase == end {
		return g.endGame()
	}
	return g.Info(), g.players[playerInTurn].Info(), err
}

func (g *Game) playData(request, origin string) dataPhase {
	c, _ := cardAction(request)
	phase := playBriscola
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	do := func(p *player.Player) (err error) {
		p.Play(c)
		g.info.PlayedCards().Add(c)
		return
	}
	nextPlayerOperator := nextPlayer
	nextPhasePredicate := g.verifyEndGame
	return dataPhase{phase, find, do, nextPlayerOperator, nextPhasePredicate}
}

func (g *Game) playEndRoundData(request, origin string) dataPhase {
	c, _ := cardAction(request)
	phase := playBriscola
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	do := func(p *player.Player) (err error) {
		p.Play(c)
		g.info.PlayedCards().Add(c)
		roundWinnerIndex := roundWinner(g)
		g.players[roundWinnerIndex].Collect(g.info.PlayedCards())
		return
	}
	nextPlayerOperator := func(uint8) uint8 {
		roundWinnerIndex := roundWinner(g)
		g.info.PlayedCards().Clear()
		return roundWinnerIndex
	}
	nextPhasePredicate := g.verifyEndGame
	return dataPhase{phase, find, do, nextPlayerOperator, nextPhasePredicate}
}

func roundWinner(g *Game) uint8 {
	return (g.playerInTurn + briscola.IndexOfWinningCard(*g.info.PlayedCards(), g.companion.Card().Seed()) + 1) % 5
}

func (g *Game) verifyEndGame() bool {
	gameHasEnded := true
	for _, pl := range g.players {
		if len(*pl.Hand()) > 0 {
			gameHasEnded = false
		}
	}
	return gameHasEnded
}

func (g *Game) endGame() ([]display.Info, []display.Info, error) {
	caller, _ := g.players.Find(notFolded)
	score1 := caller.Count() + g.companion.Ref().Count()
	score2 := uint8(0)
	for _, pl := range g.players {
		if pl != caller && pl != g.companion.Ref() {
			score2 += pl.Count()
		}
	}
	score1info := display.NewInfo("Callers", ":", strconv.Itoa(int(score1)), ";")
	score2info := display.NewInfo("Others", ":", strconv.Itoa(int(score2)), ";")
	return display.Wrap("Final Score", score1info, score2info), nil, nil
}
