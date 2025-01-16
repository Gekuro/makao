package deck_test

import (
	"makao/deck"
	"testing"
)

func TestDeck(t *testing.T) {
	mainStack := deck.NewFullCardStack()
	hand := deck.NewEmptyCardStack()

	mainStack.Shuffle()
	for range 7 {
		deck.MoveTopCard(mainStack, hand)
	}

	if len(hand.Slice()) != 7 {
		t.Errorf("expected length of hand (%d) to be 7", len(hand.Slice()))
	}

	if hand.Count() != 7 {
		t.Errorf("expected count of hand (%d) to be 7", hand.Count())
	}

	if len(mainStack.Slice()) != 45 {
		t.Errorf("expected length of main stack (%d) to be 45", len(mainStack.Slice()))
	}
	if mainStack.Count() != 45 {
		t.Errorf("expected count of main stack (%d) to be 45", mainStack.Count())
	}
}
