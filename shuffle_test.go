package texasholdempoker

import (
	"testing"
)

func Test_Shuffle(t *testing.T) {
	cards := make([]Card, len(_cards))
	copy(cards, _cards)
	t.Log("Deck before shuffle:")
	t.Log(cards)

	Shuffle(cards)

	t.Log("Deck after shuffle:")
	t.Log(cards)
}

func Test_CardsWithShuffle(t *testing.T) {
	cards := CardsWithShuffle()
	if len(cards) != 52 {
		t.Fatal("invalid cards")
	}
	t.Log(cards)
}
