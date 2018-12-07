package board

import (
	"testing"
)

func TestBoardAuctionScoreAtCreationIs0(t *testing.T) {
	if b := New(); b.AuctionScore() != 0 {
		t.Fatalf("Auction score for a new board should be 0 but is %d", b.AuctionScore())
	}
}


func TestBoardAuctionScoreCanBeSet(t *testing.T) {
	b := New()
	b.SetAuctionScore(80)
	if b.AuctionScore() != 80 {
		t.Fatalf("Auction score for a new board should be 80 but is %d", b.AuctionScore())
	}
}