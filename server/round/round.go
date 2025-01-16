package round

import (
	"fmt"
	"makao/deck"
)

type Round struct {
	// Shuffled stack of cards from which players pull cards when they cannot
	// make their turn or defend from combat cards
	pullStack *deck.CardStack

	// Never-empty stack of cards on which players place cards during turns
	placeStack *deck.CardStack

	players    []*player
	turnPlayer int
	turnState  TurnState
	// Contains integer which value has interpretation varies between turnState values
	stateInt int
}

// Creates a Makao round, with players amount denoted by first parameter
// and starting hand card count specified by the second parameter
func NewRound(playersCount int, cardCount int) *Round {
	pullStack := deck.NewFullCardStack()
	pullStack.Shuffle()

	placeStack := deck.NewEmptyCardStack()

	players := make([]*player, playersCount)
	for i := range playersCount {
		players[i] = NewPlayer()
		for range cardCount {
			deck.MoveTopCard(pullStack, players[i].hand)
		}
	}

	var rank int
	for rank > deck.TEN || rank < deck.FIVE {
		deck.MoveTopCard(pullStack, placeStack)
		_, rank, _ = placeStack.Peek()
	}

	rnd := &Round{pullStack, placeStack, players, 0, NORMAL, 0}
	return rnd
}

// Represents the action of player of provided index making their turn
// by placing cards of provided rank and suits.
//
// If rank is negative, player instead pulls card from the pull stack.
// In case of placing multiple cards the first card's suit needs to match
// the suit of the topmost card of the place stack.
//
// Returns an error if it's not the player's turn, the player doesn't
// have the specified card(s) in hand or card cannot be placed.
func (r *Round) MakeTurn(pi int, rank int, suits ...int) error {
	if r.turnPlayer != pi {
		return fmt.Errorf("it's another player's turn")
	}
	player := r.players[pi]

	if rank > 0 && (len(suits) == 0 || len(suits) > 4) {
		return fmt.Errorf("wrong suits amount (%v) for non-negative rank", len(suits))
	}

	for _, suit := range suits {
		if !player.hand.CheckForCard(suit, rank) {
			return fmt.Errorf("card of rank %v and suit %v not found in player's hand", rank, suit)
		}
	}

	// Queen resets game state by cancelling all functional cards
	if rank == deck.QUEEN {
		placeCardsAndIncrementTurn(player, r, rank, suits...)
		r.turnState, r.stateInt = NORMAL, 0
		return nil
	}

	// Check if first placed card matches place stack top card (everything matches a Queen)
	ps, pr, _ := r.placeStack.Peek()
	if pr != deck.QUEEN && suits[0] != ps && rank != pr {
		return fmt.Errorf("card of rank %v and suit %v not matching place stack top card", rank, suits[0])
	}

	switch r.turnState {
	case NORMAL:
		return makeNormalTurn(player, r, rank, suits...)
	case COMBAT:
		return makeCombatTurn(player, r, rank, suits...)
	case DEMAND:
		return makeDemandTurn(player, r, rank, suits...)
	case STOPPING:
		return makeStoppingTurn(player, r, rank, suits...)
	case SUIT_CHANGE:
		return makeNormalTurn(player, r, rank, suits...)
	default:
		return fmt.Errorf("unhandled turnState value: %v", r.turnState)
	}
}

// Refills pull stack by taking all cards below the top card from the place stack,
// shuffling them and adding them to the pull stack
func (r *Round) RefillPullStack() error {
	if r.placeStack.Count() <= 1 {
		return fmt.Errorf("not enough cards in place stack to refill pullStack")
	}
	placeSt, pullSt := deck.NewEmptyCardStack(), deck.NewEmptyCardStack()
	_ = deck.MoveTopCard(r.placeStack, placeSt)
	r.placeStack.EmptyTo(pullSt)
	r.pullStack.EmptyTo(pullSt)
	pullSt.Shuffle()

	r.placeStack, r.pullStack = placeSt, pullSt
	return nil
}

// Returns specified user's hand as array of strings.
// Returns nil error only if user exists
func (r *Round) GetPlayerHand(idx int) ([]string, error) {
	if len(r.players) <= idx {
		return nil, fmt.Errorf("wrong player index")
	}
	return r.players[idx].hand.Slice(), nil
}

// Returns first (and only visible) card of the place stack in string format
func (r *Round) GetTopCard() string {
	str, _ := r.placeStack.PeekStr()
	return str
}
