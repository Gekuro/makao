package round

type TurnState int

// TODO define stateInt values for each state
const (
	// During NORMAL state the current player has to give a card which
	// matches the top card on the place stack in suit or rank
	NORMAL TurnState = iota

	// During COMBAT state the current player has to either give a combat card
	// (cards of rank 2, 3, K♥️ or K♠️) which matches last card in suit or rank,
	// or pull the sum of the previous combat cards (combat Kings count as 5)
	COMBAT

	// During DEMAND state the current player has to give the non-functional
	// card with the rank chosen by the player which started the demand (placed a Jack)
	DEMAND

	// During STOPPING state the current player has to defend by placing a 4
	// or stop for amount of rounds equal to the amount of 4 which were placed
	STOPPING

	// During SUIT_CHANGE state the top card of the place stack is an Ace, but can only
	// be matched by rank or suit which was selected by the person which placed the card
	SUIT_CHANGE
)

// Returns string representation of TurnState value
func (t TurnState) GetTurnStateStr() string {
	var str string
	switch t {
	case NORMAL:
		str = "NORMAL"
	case COMBAT:
		str = "COMBAT"
	case DEMAND:
		str = "DEMAND"
	case STOPPING:
		str = "STOPPING"
	case SUIT_CHANGE:
		str = "SUIT_CHANGE"
	}
	return str
}
