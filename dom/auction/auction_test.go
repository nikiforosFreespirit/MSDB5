package auction

import (
	"strconv"
	"testing"
)

type auctionTest struct {
	auction Score
}

func (testObject *auctionTest) set(value Score) {
	testObject.auction = value
}
func (testObject *auctionTest) get() Score {
	return testObject.auction
}

var initialValue Score
var minValue Score = 61
var maxValue Score = 120

func TestRaiseAuctionScoreFirstAssignmentShouldBeSuperiorThan61ElseEither61(t *testing.T) {
	const currentValue = 1
	testObject := auctionTest{initialValue}
	initialValue.Update(currentValue, testObject.set)
	testPlayerScore(t, testObject.auction, minValue)
}

func TestInvalidRaiseAuctionScoreFirstAssignmentShouldBeAlways61(t *testing.T) {
	const currentValue = 0
	testObject := auctionTest{initialValue}
	initialValue.Update(currentValue, testObject.set)
	testPlayerScore(t, testObject.auction, minValue)
}

func TestRaiseAuctionTo65(t *testing.T) {
	const currentValue = 65
	testObject := auctionTest{initialValue}
	initialValue.Update(currentValue, testObject.set)
	testPlayerScore(t, testObject.auction, currentValue)
}
func TestRaiseAuctionTo135ShouldStopAt120(t *testing.T) {
	const currentValue = 135
	testObject := auctionTest{initialValue}
	initialValue.Update(currentValue, testObject.set)
	testPlayerScore(t, testObject.auction, maxValue)
}

func TestPlayerRaisingAuctionAfterAnotherWithLowerScore(t *testing.T) {
	value1 := Score(94)
	testObject := auctionTest{value1}
	const value2 = 90
	value1.Update(value2, testObject.set)
	testPlayerScore(t, testObject.auction, value1)
}

func Test2PlayersRaisingAuction(t *testing.T) {
	value1 := Score(65)
	const value2 = 80
	testObject := auctionTest{initialValue}
	initialValue.Update(value1, testObject.set)
	value1.Update(value2, testObject.set)
	testPlayerScore(t, testObject.auction, value2)
}

func TestCheckAndUpdate_OK(t *testing.T) {
	const value = 80
	testObject := auctionTest{0}
	CheckAndUpdate(strconv.Itoa(value), func() bool { return false }, func() {}, testObject.get, testObject.set)
	testPlayerScore(t, testObject.auction, value)
}

func TestCheckAndUpdate_Fold(t *testing.T) {
	testObject := auctionTest{initialValue}
	CheckAndUpdate("ciao", func() bool { return false }, func() {}, testObject.get, testObject.set)
	testPlayerScore(t, testObject.auction, initialValue)
}

func testPlayerScore(t *testing.T, actualScore, expectedScore Score) {
	if actualScore != expectedScore {
		t.Fatalf("Auction score should be set at %d but is %d", expectedScore, actualScore)
	}
}
