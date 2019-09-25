package game

import (
	"fmt"
	"strings"

	"github.com/mcaci/ita-cards/card"
)

type Round struct {
	*Game
	req  string
	rErr error
}

func (g Round) Card() (*card.Item, error) {
	fields := strings.Split(g.req, "#")
	if len(fields) > 2 {
		return card.New(fields[1], fields[2])
	}
	return nil, fmt.Errorf("not enough data to make a card: %s", g.req)
}
func (g Round) Value() string     { return parse(g.req, val) }
func (g Round) RoundError() error { return g.rErr }
func (g Round) PlayedCard() card.Item {
	c, err := g.Card()
	if err != nil {
		return card.Item{}
	}
	return *c
}
