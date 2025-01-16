package deck

import "fmt"

type card struct {
	suit int
	rank int
	next *card
}

func (c *card) String() string {
	return fmt.Sprintf("%s%s", Ranks[c.rank], Suits[c.suit])
}

func (c *card) Suit() int {
	return c.suit
}

func (c *card) Rank() int {
	return c.rank
}
