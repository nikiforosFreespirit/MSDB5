package play

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/app/action"
	"github.com/nikiforosFreespirit/msdb5/dom/briscola"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/playerset"
)

type PlayCardStruct struct {
	request, origin string
	players         playerset.Players
	playedCards     *deck.Cards
	sideDeck        *deck.Cards
	briscolaSeed    card.Seed
}

func NewPlay(request, origin string, players playerset.Players,
	playedCards *deck.Cards, sideDeck *deck.Cards, briscolaSeed card.Seed) action.Executer {
	return &PlayCardStruct{request, origin, players,
		playedCards, sideDeck, briscolaSeed}
}

var playersRoundRobin = func(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }

func (pcs PlayCardStruct) Do(p *player.Player) error {
	data := strings.Split(pcs.request, "#")
	number := data[1]
	seed := data[2]
	c, err := card.Create(number, seed)
	err = p.Play(c)
	if err != nil {
		return err
	}
	pcs.playedCards.Add(c)
	roundHasEnded := len(*pcs.playedCards) == 5
	if roundHasEnded {
		playerInTurn, _ := pcs.players.FindIndex(func(pl *player.Player) bool { return pl == p })
		winningCardIndex := briscola.IndexOfWinningCard(*pcs.playedCards, pcs.briscolaSeed)
		next := playersRoundRobin(uint8(playerInTurn) + winningCardIndex)
		pcs.players[next].Collect(pcs.playedCards)
		var isHandEmpty = func(p *player.Player) bool { return p.IsHandEmpty() }
		if pcs.players.All(isHandEmpty) && len(*pcs.sideDeck) > 0 {
			pcs.players[next].Collect(pcs.sideDeck)
			pcs.sideDeck.Clear()
		}
	}
	return err
}
