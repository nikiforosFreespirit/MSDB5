package notify

import (
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func createInGameMsg(gameInfo selfInformer, pl *player.Player) string {
	printer := message.NewPrinter(language.English)
	me := printer.Sprintf("Player: %+v\n", pl)
	if gameInfo.Phase() == phase.ExchangingCards {
		me += printer.Sprintf("Side deck: %+v\n", *gameInfo.SideDeck())
	}
	return me
}

func createSideGameMsg(gameInfo sidedeckInformer, quantity uint8) string {
	printer := message.NewPrinter(language.English)
	return printer.Sprintf("Side deck section: %+v\n", (*gameInfo.SideDeck())[:quantity])
}

func createMlMsg(gameInfo miner) (bool, string) {
	var msg string
	canLog := true
	printer := message.NewPrinter(language.English)
	switch gameInfo.Phase() {
	case phase.ChoosingCompanion:
		msg = printer.Sprintf("%s, %s, %d\n", gameInfo.CurrentPlayer().Name(), gameInfo.Companion().Name(), *(gameInfo.AuctionScore()))
	case phase.PlayingCards:
		msg = printer.Sprintf("%s, %d\n", gameInfo.CurrentPlayer().Name(), gameInfo.LastCardPlayed())
	case phase.End:
		msg = printer.Sprintf("%s\n", gameInfo.CurrentPlayer().Name())
	default:
		canLog = false
	}
	return canLog, msg
}
