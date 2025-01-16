package round

import "makao/deck"

type player struct {
	hand *deck.CardStack
}

// Returns a new player with an empty hand
func NewPlayer() *player {
	return &player{deck.NewEmptyCardStack()}
}
