package round

import "encoding/json"

type RoundSnapshot struct {
	TopCard        string     `json:"topCard"`
	PullStackCount int        `json:"pullStackCount"`
	Players        [][]string `json:"playerHands"`
	TurnPlayerIdx  int        `json:"turnPlayerIdx"`
	TurnState      string     `json:"turnState"`
	StateValue     int        `json:"stateValueCounter"`
}

// Returns RoundSnapshot object, containing de-encapsulated data
func (r *Round) GetRoundSnapshot() *RoundSnapshot {
	playerHands := make([][]string, len(r.players))
	for i := range r.players {
		playerHands[i] = r.players[i].hand.Slice()
	}

	return &RoundSnapshot{
		r.GetTopCard(),
		r.pullStack.Count(),
		playerHands,
		r.turnPlayer,
		r.turnState.GetTurnStateStr(),
		r.stateInt,
	}
}

// Returns JSON round snapshot for logging purposes.
// Contains data which should be hidden from players
func (r *Round) GetRoundJSON() []byte {
	snap := r.GetRoundSnapshot()
	payload, err := json.Marshal(snap)
	if err != nil {
		return []byte("{\"error\":\"round data corrupted\"}")
	}
	return payload
}
