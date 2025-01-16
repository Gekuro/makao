package deck

import (
	"fmt"
	"math/rand/v2"
)

// Moves top card from source stack to the top of destination stack.
// Returns error if source stack is empty
func MoveTopCard(src *CardStack, dst *CardStack) error {
	if src.count == 0 {
		return fmt.Errorf("not enough cards in stack")
	}

	card := src.top
	src.top = card.next
	card.next = dst.top
	dst.top = card

	src.count--
	dst.count++
	return nil
}

// Finds card with the provided suit and rank in the source stack
// then moves it to the top of the destination stack.
// Returns non-nil error if source stack is empty or doesn't contain the card
func MoveCardBySuitAndRank(suit int, rank int, src *CardStack, dst *CardStack) error {
	if src.count == 0 {
		return fmt.Errorf("not enough cards in stack")
	}
	if src.top.suit == suit && src.top.rank == rank {
		err := MoveTopCard(src, dst)
		if err != nil {
			return err
		}
		return nil
	}

	prev := src.top
	cursor := src.top.next
	for {
		if cursor == nil {
			return fmt.Errorf("card not in stack")
		}
		if cursor.suit == suit && cursor.rank == rank {
			prev.next = cursor.next
			cursor.next = dst.top
			dst.top = cursor

			src.count--
			dst.count++
			return nil
		}
		cursor = cursor.next
	}
}

func (c *CardStack) Shuffle() {
	if c.count == 0 {
		return
	}

	var s []*card = make([]*card, c.count)
	cursor := c.top
	for i := range s {
		s[i] = cursor
		cursor = cursor.next
	}

	rand.Shuffle(c.count, func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	cursor = s[0]
	c.top = cursor
	for i := 1; i < c.count; i++ {
		cursor.next = s[i]
		cursor = cursor.next
	}
}

func (c *CardStack) EmptyTo(dst *CardStack) {
	for {
		err := MoveTopCard(c, dst)
		if err != nil {
			return
		}
	}
}
