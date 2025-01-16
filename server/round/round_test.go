package round_test

import (
	"makao/deck"
	"makao/round"
	"testing"
)

func NormalTurnTest(t *testing.T) {
	decks := [][][2]int{
		[][2]int{
			[2]int{deck.DIAMONDS, deck.FIVE},
			[2]int{deck.HEARTS, deck.KING},
			[2]int{deck.DIAMONDS, deck.THREE},
		},
		[][2]int{
			[2]int{deck.HEARTS, deck.THREE},
			[2]int{deck.HEARTS, deck.TWO},
			[2]int{deck.SPADES, deck.FIVE},
		},
	}
	r, err := round.MockRound(deck.DIAMONDS, deck.SEVEN, decks)
	if err != nil {
		t.Fatal("impossible test data, make sure all provided mock round cards are unique")
	}
}
