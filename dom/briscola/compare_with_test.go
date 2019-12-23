package briscola

import "testing"
import "github.com/mcaci/ita-cards/card"

type testSeeder card.Seed

func (t testSeeder) Seed() card.Seed {
	return card.Seed(t)
}

func TestScenarioWithAceOfCoinWinning(t *testing.T) {
	// testing 1 and 2 of Coin, briscola is Coin
	verifyRoundScenario(t, 1, 2, testSeeder(card.Coin), false)
}

func TestScenarioWithTwoOfCoinLosing(t *testing.T) {
	// 2 and 3 of Coin, briscola is Coin
	verifyRoundScenario(t, 2, 3, testSeeder(card.Coin), true)
}

func TestScenarioWithSixOfCoinWinningBecauseHigher(t *testing.T) {
	// 5 and 6 of Coin, briscola is Coin
	verifyRoundScenario(t, 5, 6, testSeeder(card.Coin), true)
}

func TestScenarioWithSixOfCoinWinningBecausePlayedFirst(t *testing.T) {
	// 6 and 5 of Coin, briscola is Coin
	verifyRoundScenario(t, 6, 5, testSeeder(card.Coin), false)
}

func TestScenarioWithTenOfCoinWinning(t *testing.T) {
	// 10 and 4 of Coin, briscola is Cup
	verifyRoundScenario(t, 10, 4, testSeeder(card.Cup), false)
}

func TestScenarioWithTenOfCoinLosing(t *testing.T) {
	// 10 and 4 of Coin, briscola is Cup
	verifyRoundScenario(t, 10, 3, testSeeder(card.Coin), true)
}

func TestScenarioWithTwoOfSwordsWinningBecauseOfBriscola(t *testing.T) {
	// 3 of Coin and 2 of Sword, briscola is Sword
	verifyRoundScenario(t, 3, 22, testSeeder(card.Sword), true)
}

func verifyRoundScenario(t *testing.T, a, b uint8, briscola testSeeder, expectedWinner bool) {
	if index := doesOtherCardWin(*card.MustID(a), *card.MustID(b), briscola); index != expectedWinner {
		t.Fatal("Unexpected winner")
	}
}