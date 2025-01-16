package deck

var Suits []string = []string{"♥️", "♠️", "♦️", "♣️"}
var Ranks []string = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

const (
	HEARTS = iota
	SPADES
	DIAMONDS
	CLUBS
)

const (
	ACE = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

type CardStack struct {
	top   *card
	count int
}

func NewEmptyCardStack() *CardStack {
	return &CardStack{nil, 0}
}

// Returns pointer of a new CardStack containing 52 cards.
// Cards are sorted first by suit in order: "♥️", "♠️", "♦️", "♣️"
// then by rank in descending order
func NewFullCardStack() *CardStack {
	stc := NewEmptyCardStack()
	for sx := range Suits {
		for rx := range Ranks {
			c := card{sx, rx, stc.top}
			stc.top = &c
			stc.count++
		}
	}
	return stc
}

// Returns suit and rank of top card.
// If stack is empty returns 0, 0 and false
func (c *CardStack) Peek() (suit int, rank int, ok bool) {
	if c.count == 0 {
		return 0, 0, false
	}
	return c.top.suit, c.top.rank, true
}

// Returns top card value in string format, example: "10♠️".
// If stack is empty returns empty string and false boolean value
func (c *CardStack) PeekStr() (string, bool) {
	if c.count == 0 {
		return "", false
	}
	return c.top.String(), true
}

// Returns true if card is in stack, and false otherwise
func (c *CardStack) CheckForCard(suit int, rank int) bool {
	cursor := c.top
	for cursor != nil {
		if cursor.suit == suit && cursor.rank == rank {
			return true
		}
		cursor = cursor.next
	}
	return false
}

func (c *CardStack) Count() int {
	return c.count
}

// Creates a slice of strings from CardStack
func (c *CardStack) StringSlice() []string {
	s := make([]string, c.count)
	cursor := c.top
	for i := range s {
		s[i] = cursor.String()
		cursor = cursor.next
	}
	return s
}

// Creates a slice of cards from CardStack
func (c *CardStack) Slice() []card {
	s := make([]card, c.count)
	cursor := c.top
	for i := range s {
		s[i] = *cursor
		cursor = cursor.next
	}
	return s
}
