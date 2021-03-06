package game

import (
	"container/list"
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/game/start"
	"github.com/mcaci/msdb5/app/game/track"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

// Game struct
type Game struct {
	lastPlaying       list.List
	players           team.Players
	caller, companion *player.Player
	briscolaCard      card.Item
	side, playedCards set.Cards
	auctionScore      auction.Score
	phase             phase.ID
}

// NewGame func
func NewGame(withSide bool) *Game {
	g := new(Game)
	start.Players(&g.players)
	start.DistributeCards(g, withSide)
	track.Player(&g.lastPlaying, g.players[0])
	return g
}

func (g *Game) AuctionScore() *auction.Score  { return &g.auctionScore }
func (g *Game) Briscola() card.Item           { return g.briscolaCard }
func (g *Game) Caller() *player.Player        { return g.caller }
func (g *Game) Companion() *player.Player     { return g.companion }
func (g *Game) CurrentPlayer() *player.Player { return g.lastPlaying.Front().Value.(*player.Player) }
func (g *Game) IsSideUsed() bool              { return len(g.side) > 0 }
func (g *Game) LastPlayer() *player.Player    { return g.lastPlaying.Back().Value.(*player.Player) }
func (g *Game) LastPlaying() *list.List       { return &g.lastPlaying }
func (g *Game) Phase() phase.ID               { return g.phase }
func (g *Game) Players() team.Players         { return g.players }
func (g *Game) PlayedCards() *set.Cards       { return &g.playedCards }
func (g *Game) SideDeck() *set.Cards          { return &g.side }
func (g *Game) IsSideToShow() bool            { return len(g.SideSubset()) > 0 && g.phase == phase.InsideAuction }
func (g *Game) SideSubset() set.Cards {
	var quantity uint8
	if g.IsSideUsed() {
		newScore := uint8(*g.AuctionScore())
		quantity = uint8(newScore/90 + newScore/100 + newScore/110 + newScore/120 + newScore/120)
	}
	return g.side[:quantity]
}

func (g *Game) SetAuction(s auction.Score)      { g.auctionScore = s }
func (g *Game) SetBriscola(c *card.Item)        { g.briscolaCard = *c }
func (g *Game) SetCaller(pred player.Predicate) { g.caller = g.players.At(g.players.MustIndex(pred)) }
func (g *Game) SetCompanion(pl *player.Player)  { g.companion = pl }
func (g *Game) setPhase(ph phase.ID)            { g.phase = ph }

func (g Game) String() string {
	return fmt.Sprintf("(Turn of: %s, Companion is: %s, Played cards: %v, Auction score: %d, Phase: %s, Players: %v, Side Deck: %v, Last Players: %v)",
		g.CurrentPlayer().Name(), g.briscolaCard, g.playedCards, g.auctionScore, g.phase, g.players, g.side, g.lastPlaying)
}
