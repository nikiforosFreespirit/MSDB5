package msg

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/msdb5/app/input"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/score"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/team"
)

func toOS(g roundInformer, inputRequest, origin string) {
	io.WriteString(os.Stdout, inputRequest)

	rErr := g.RoundError()
	if rErr != nil {
		errMsg := fmtErr(g, inputRequest, rErr)
		io.WriteString(os.Stdout, errMsg)
		return
	}
	s := team.Sender(senderInfo{g.Players(), origin})
	senderInfo := fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", s.Name(), inputRequest, s, g)
	io.WriteString(os.Stdout, senderInfo)

	// compute score
	pilers := make([]score.Piler, len(g.Players()))
	for i, p := range g.Players() {
		pilers[i] = p
	}
	scoreTeam1, scoreTeam2 := score.Calc(g.Caller(), g.Companion(), pilers, briscola.Points)
	scoreMsg := fmt.Sprintf("Scores -> Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
	io.WriteString(os.Stdout, scoreMsg)
}

func fmtErr(g roundInformer, inputRequest string, rErr error) string {
	errMsg := fmt.Sprintf("Error: %+v\n", rErr)
	if rErr == phase.ErrUnexpectedPhase {
		_, id := phase.ToID(input.Value(inputRequest))
		errMsg = fmt.Sprintf("Phase is not %d but %d", id, g.Phase())
	}
	if rErr == team.ErrUnexpectedPlayer {
		errMsg = fmt.Sprintf("Expecting player %s to play", g.CurrentPlayer().Name())
	}
	return errMsg
}