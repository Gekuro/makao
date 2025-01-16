package round

import (
	"fmt"
	"makao/deck"
)

// Mocks a makao round from provided suit and rank of top place stack card and player hands.
// Hands are expressed as arrays of integer tuples where first int is suit and second is rank.
func MockRound(topSuit int, topRank int, hands [][][2]int) (*Round, error) {
	pullStack := deck.NewFullCardStack()
	pullStack.Shuffle()

	placeStack := deck.NewEmptyCardStack()

	players := make([]*player, len(hands))
	for i, hand := range hands {
		players[i] = NewPlayer()
		for _, cardTuple := range hand {
			if deck.MoveCardBySuitAndRank(cardTuple[0], cardTuple[1], pullStack, players[i].hand) != nil {
				return nil, fmt.Errorf("found duplicate cards in provided card data")
			}
		}
	}

	if deck.MoveCardBySuitAndRank(topSuit, topRank, pullStack, placeStack) != nil {
		return nil, fmt.Errorf("found duplicate cards in provided card data")
	}

	rnd := &Round{pullStack, placeStack, players, 0, NORMAL, 0}
	return rnd, nil
}
