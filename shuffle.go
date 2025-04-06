package texasholdempoker

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

var _ rand.Source = (*CrytoRandomSource)(nil)

type CrytoRandomSource struct{}

func NewCryptoRandomSource() rand.Source {
	return &CrytoRandomSource{}
}

func (c *CrytoRandomSource) Int63() int64 {
	var b [8]byte

	crand.Read(b[:])

	return int64(binary.LittleEndian.Uint64(b[:])) & (1<<63 - 1)
}

func (c *CrytoRandomSource) Seed(seed int64) {}

var r = rand.New(NewCryptoRandomSource())

// 洗牌
func Shuffle[T Card | int16](cards []T) {
	r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

// 生成一副52张洗好且类型为int16的牌
func FullCardsWithShuffle() []int16 {
	cards := make([]int16, len(_fullcards))
	copy(cards, _fullcards)
	Shuffle(cards)
	return cards
}

// 生成一副52张洗好且类型为Card的牌
func CardsWithShuffle() []Card {
	cards := make([]Card, len(_cards))
	copy(cards, _cards)
	Shuffle(cards)
	return cards
}
