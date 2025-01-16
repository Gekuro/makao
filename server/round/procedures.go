package round

import (
	"fmt"
	"makao/deck"
	"slices"
)

func makeNormalTurn(pl *player, r *Round, rank int, suits ...int) error {
	switch {
	case rank == -1: // Player is pulling a card
		pullCardsAndIncrementTurn(pl, r, 1)
		return nil
	case rank == deck.TWO || rank == deck.THREE: // Combat cards
		r.turnState, r.stateInt = COMBAT, 0
		return makeCombatTurn(pl, r, rank, suits...)
	case rank == deck.FOUR: // Stopping card
		return nil
	case rank == deck.KING: // King (either combat or non-functional card)
		if slices.Contains(suits, 0) || slices.Contains(suits, 1) {
			r.turnState, r.stateInt = COMBAT, 0
			return makeCombatTurn(pl, r, rank, suits...)
		} else {
			placeCardsAndIncrementTurn(pl, r, rank, suits...)
			return nil
		}
	default: // Non-functional cards
		placeCardsAndIncrementTurn(pl, r, rank, suits...)
		return nil
	}

}

func makeCombatTurn(pl *player, r *Round, rank int, suits ...int) error {
	// If player concedes the turn
	if rank == -1 {
		pullCardsAndIncrementTurn(pl, r, r.stateInt)
		r.turnState, r.stateInt = NORMAL, 0
		return nil
	}

	if rank == deck.TWO || rank == deck.THREE {
		placeCardsAndIncrementTurn(pl, r, rank, suits...)
		r.stateInt += (rank + 1) * len(suits)
		return nil
	}

	// Each functional king is worth 5 cards (to be pulled)
	var cost int
	if slices.Contains(suits, 0) {
		cost += 5
	}
	if slices.Contains(suits, 1) {
		cost += 5
	}
	if cost == 0 {
		return fmt.Errorf("cannot play only non-functional Kings during combat")
	}

	placeCardsAndIncrementTurn(pl, r, rank, suits...)
	r.stateInt += cost
	return nil
}

func makeDemandTurn(pl *player, r *Round, rank int, suits ...int) error {
	return nil
}

func makeStoppingTurn(pl *player, r *Round, rank int, suits ...int) error {
	return nil
}

func makeSuitChangeTurn(pl *player, r *Round, rank int, suits ...int) error {
	return nil
}

// Represents a player placing cards on the place stack and ending the turn.
// Assumes the move and provided suits and ranks are valid, and the player
// has the cards in hand
func placeCardsAndIncrementTurn(pl *player, r *Round, rank int, suits ...int) {
	for _, suit := range suits {
		_ = deck.MoveCardBySuitAndRank(suit, rank, pl.hand, r.placeStack)
	}
	r.turnPlayer = (r.turnPlayer + 1) % len(r.players)
}

// Represents a player pulling cards and ending the turn.
// If refilling pull stack is impossible, simply increments turn
func pullCardsAndIncrementTurn(pl *player, r *Round, amount int) {
	if r.pullStack.Count() <= amount {
		_ = r.RefillPullStack()
	}
	for range amount {
		_ = deck.MoveTopCard(r.pullStack, pl.hand)
	}
	r.turnPlayer = (r.turnPlayer + 1) % len(r.players)
}
