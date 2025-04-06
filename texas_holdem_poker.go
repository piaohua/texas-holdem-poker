package texasholdempoker

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	NoPair        int8 = 1  // 高牌
	OnePair       int8 = 2  // 对子
	TwoPair       int8 = 3  // 两对
	ThreeOfAKind  int8 = 4  // 三条
	Straight      int8 = 5  // 顺子
	Flush         int8 = 6  // 同花
	FullHouse     int8 = 7  // 葫芦
	FourOfAKind   int8 = 9  // 四条（金刚）
	StraightFlush int8 = 10 // 同花顺
	RoyalFlush    int8 = 11 // 皇家同花顺
)

const (
	SPADE   int8 = 4 // 黑桃
	HEART   int8 = 3 // 红桃
	CLUB    int8 = 2 // 梅花
	DIAMOND int8 = 1 // 方块

	spade   = "♠️"
	heart   = "♥️"
	club    = "♣️"
	diamond = "♦️"
)

type Card struct {
	Color int8
	Value int8
}

func (c Card) String() string {
	b := strings.Builder{}
	b.Grow(3)
	switch c.Value {
	case 0x0d:
		b.WriteString(" K")
	case 0x0c:
		b.WriteString(" Q")
	case 0x0b:
		b.WriteString(" J")
	case 0x0e:
		b.WriteString(" A")
	default:
		b.WriteString(fmt.Sprintf("%2d", c.Value))
	}
	switch c.Color {
	case SPADE:
		b.WriteString(spade)
	case HEART:
		b.WriteString(heart)
	case CLUB:
		b.WriteString(club)
	case DIAMOND:
		b.WriteString(diamond)
	}
	return b.String()
}

var (
	_fullcards = []int16{
		0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, // 黑桃
		0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, // 红桃
		0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, // 梅花
		0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, // 方块
	}

	_cards = []Card{
		{4, 2}, {4, 3}, {4, 4}, {4, 5}, {4, 6}, {4, 7}, {4, 8}, {4, 9}, {4, 0xa}, {4, 0xb}, {4, 0xc}, {4, 0xd}, {4, 0xe},
		{3, 2}, {3, 3}, {3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8}, {3, 9}, {3, 0xa}, {3, 0xb}, {3, 0xc}, {3, 0xd}, {3, 0xe},
		{2, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {2, 8}, {2, 9}, {2, 0xa}, {2, 0xb}, {2, 0xc}, {2, 0xd}, {2, 0xe},
		{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}, {1, 8}, {1, 9}, {1, 0xa}, {1, 0xb}, {1, 0xc}, {1, 0xd}, {1, 0xe},
	}

	// 0x41 >> 4   == color
	// 0x41 & 0x0f == value
	_setCARDS = map[int16]Card{
		0x42: {4, 2},
		0x43: {4, 3},
		0x44: {4, 4},
		0x45: {4, 5},
		0x46: {4, 6},
		0x47: {4, 7},
		0x48: {4, 8},
		0x49: {4, 9},
		0x4a: {4, 10},
		0x4b: {4, 11},
		0x4c: {4, 12},
		0x4d: {4, 13},
		0x4e: {4, 14},
		0x32: {3, 2},
		0x33: {3, 3},
		0x34: {3, 4},
		0x35: {3, 5},
		0x36: {3, 6},
		0x37: {3, 7},
		0x38: {3, 8},
		0x39: {3, 9},
		0x3a: {3, 10},
		0x3b: {3, 11},
		0x3c: {3, 12},
		0x3d: {3, 13},
		0x3e: {3, 14},
		0x22: {2, 2},
		0x23: {2, 3},
		0x24: {2, 4},
		0x25: {2, 5},
		0x26: {2, 6},
		0x27: {2, 7},
		0x28: {2, 8},
		0x29: {2, 9},
		0x2a: {2, 10},
		0x2b: {2, 11},
		0x2c: {2, 12},
		0x2d: {2, 13},
		0x2e: {2, 14},
		0x12: {1, 2},
		0x13: {1, 3},
		0x14: {1, 4},
		0x15: {1, 5},
		0x16: {1, 6},
		0x17: {1, 7},
		0x18: {1, 8},
		0x19: {1, 9},
		0x1a: {1, 10},
		0x1b: {1, 11},
		0x1c: {1, 12},
		0x1d: {1, 13},
		0x1e: {1, 14},
	}

	// 转牌6取5组合
	// _turnCombinations = [6][5]int{
	// 	{0, 1, 2, 3, 4},
	// 	{0, 1, 2, 3, 5},
	// 	{0, 1, 2, 4, 5},
	// 	{0, 1, 3, 4, 5},
	// 	{0, 2, 3, 4, 5},
	// 	{1, 2, 3, 4, 5},
	// }

	// 河牌7取5组合
	_riverCombinations = [21][5]int{
		{0, 1, 2, 3, 4},
		{0, 1, 2, 3, 5},
		{0, 1, 2, 3, 6},
		{0, 1, 2, 4, 5},
		{0, 1, 2, 4, 6},
		{0, 1, 2, 5, 6},
		{0, 1, 3, 4, 5},
		{0, 1, 3, 4, 6},
		{0, 1, 3, 5, 6},
		{0, 1, 4, 5, 6},
		{0, 2, 3, 4, 5},
		{0, 2, 3, 4, 6},
		{0, 2, 3, 5, 6},
		{0, 2, 4, 5, 6},
		{0, 3, 4, 5, 6},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 6},
		{1, 2, 3, 5, 6},
		{1, 2, 4, 5, 6},
		{1, 3, 4, 5, 6},
		{2, 3, 4, 5, 6},
	}

	borad5Pool = sync.Pool{
		New: func() any {
			return new(BoradCards)
		},
	}
	borad7Pool = sync.Pool{
		New: func() any {
			return new([7]Card)
		},
	}
)

type HoleCards [2]Card  // 手牌
type BoradCards [5]Card // 牌面

func (b BoradCards) String() string {
	s := strings.Builder{}
	for i := range b {
		s.WriteString(fmt.Sprintf(" %3s |", b[i]))
	}
	return s.String()
}

// 是否是同花
func (b *BoradCards) isFlush() bool {
	return b[0].Color == b[1].Color &&
		b[0].Color == b[2].Color &&
		b[0].Color == b[3].Color &&
		b[0].Color == b[4].Color
}

// 皇家顺子(忽略花色)
func (b *BoradCards) isRoyalStraight() bool {
	return b[0].Value == 0x0a &&
		b[1].Value == 0x0b &&
		b[2].Value == 0x0c &&
		b[3].Value == 0x0d &&
		b[4].Value == 0x0e
}

// 金刚
func (b *BoradCards) isFourOfAKind() bool {
	// 1111x
	if b[1].Value == b[0].Value &&
		b[2].Value == b[0].Value &&
		b[3].Value == b[0].Value &&
		b[4].Value != b[0].Value {
		return true
	}
	// x1111
	if b[0].Value != b[1].Value &&
		b[2].Value == b[1].Value &&
		b[3].Value == b[1].Value &&
		b[4].Value == b[1].Value {
		return true
	}
	return false
}

// 葫芦
func (b *BoradCards) isFullHouse() bool {
	// 111xx
	if b[0].Value == b[1].Value &&
		b[0].Value == b[2].Value &&
		b[0].Value != b[3].Value &&
		b[3].Value == b[4].Value {
		return true
	}
	// xx111
	if b[0].Value == b[1].Value &&
		b[0].Value != b[2].Value &&
		b[2].Value == b[3].Value &&
		b[2].Value == b[4].Value {
		return true
	}
	return false
}

// 顺子(忽略花色)
func (b *BoradCards) isStraight() bool {
	if b[0].Value == 0x02 &&
		b[1].Value == 0x03 &&
		b[2].Value == 0x04 &&
		b[3].Value == 0x05 &&
		b[4].Value == 0x0e {
		return true
	}
	return (b[0].Value+1) == b[1].Value &&
		(b[1].Value+1) == b[2].Value &&
		(b[2].Value+1) == b[3].Value &&
		(b[3].Value+1) == b[4].Value
}

// 三条
func (b *BoradCards) isThreeOfAKind() bool {
	// 111xx
	if b[0].Value == b[1].Value &&
		b[0].Value == b[2].Value &&
		b[0].Value != b[3].Value &&
		b[0].Value != b[4].Value &&
		b[3].Value != b[4].Value {
		return true
	}
	// x111x
	if b[1].Value == b[2].Value &&
		b[1].Value == b[3].Value &&
		b[1].Value != b[0].Value &&
		b[1].Value != b[4].Value &&
		b[0].Value != b[4].Value {
		return true
	}
	// xx111
	if b[2].Value == b[3].Value &&
		b[2].Value == b[4].Value &&
		b[2].Value != b[0].Value &&
		b[2].Value != b[1].Value &&
		b[0].Value != b[1].Value {
		return true
	}
	return false
}

// 两对
func (b *BoradCards) isTwoPair() bool {
	// 1122x
	if b[0].Value == b[1].Value &&
		b[2].Value == b[3].Value &&
		b[0].Value != b[2].Value &&
		b[0].Value != b[4].Value &&
		b[2].Value != b[4].Value {
		return true
	}
	// 11x33
	if b[0].Value == b[1].Value &&
		b[4].Value == b[3].Value &&
		b[0].Value != b[4].Value &&
		b[0].Value != b[2].Value &&
		b[2].Value != b[4].Value {
		return true
	}
	// x2233
	if b[1].Value == b[2].Value &&
		b[4].Value == b[3].Value &&
		b[0].Value != b[4].Value &&
		b[0].Value != b[2].Value &&
		b[2].Value != b[4].Value {
		return true
	}
	return false
}

// 对子
func (b *BoradCards) isOnePair() bool {
	// 11xxx
	if b[0].Value == b[1].Value &&
		b[0].Value != b[2].Value &&
		b[0].Value != b[3].Value &&
		b[0].Value != b[4].Value &&
		b[2].Value != b[3].Value &&
		b[2].Value != b[4].Value &&
		b[3].Value != b[4].Value {
		return true
	}
	// x11xx
	if b[1].Value == b[2].Value &&
		b[1].Value != b[0].Value &&
		b[1].Value != b[3].Value &&
		b[1].Value != b[4].Value &&
		b[0].Value != b[3].Value &&
		b[0].Value != b[4].Value &&
		b[3].Value != b[4].Value {
		return true
	}
	// xx11x
	if b[2].Value == b[3].Value &&
		b[2].Value != b[0].Value &&
		b[2].Value != b[1].Value &&
		b[2].Value != b[4].Value &&
		b[0].Value != b[1].Value &&
		b[0].Value != b[4].Value &&
		b[1].Value != b[4].Value {
		return true
	}
	// xxx11
	if b[3].Value == b[4].Value &&
		b[3].Value != b[0].Value &&
		b[3].Value != b[1].Value &&
		b[3].Value != b[2].Value &&
		b[0].Value != b[1].Value &&
		b[0].Value != b[2].Value &&
		b[1].Value != b[2].Value {
		return true
	}
	return false
}

// 计算当前手牌b牌力是否大于等于v，返回b牌力值，牌力小于v时返回0
func (b *BoradCards) getValue(v int8) (value int8) {
	switch v {
	case RoyalFlush: // 皇家同花顺
		return
	case StraightFlush: // 同花顺
		if b.isFlush() {
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		}
		return
	case FourOfAKind: // 四条（金刚）
		if b.isFlush() {
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		} else if b.isFourOfAKind() {
			value = FourOfAKind
		}
		return
	case FullHouse: // 葫芦
		if b.isFlush() {
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		} else if b.isFourOfAKind() {
			value = FourOfAKind
		} else if b.isFullHouse() {
			value = FullHouse
		}
		return
	case Flush: // 同花
		if b.isFlush() {
			value = Flush
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		} else if b.isFourOfAKind() {
			value = FourOfAKind
		} else if b.isFullHouse() {
			value = FullHouse
		}
		return
	case Straight: // 顺子
		if b.isFlush() {
			value = Flush
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		} else if b.isFourOfAKind() {
			value = FourOfAKind
		} else if b.isFullHouse() {
			value = FullHouse
		} else if b.isRoyalStraight() {
			value = Straight
		} else if b.isStraight() {
			value = Straight
		}
		return
	case ThreeOfAKind: // 三条
		if b.isFlush() {
			value = Flush
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		} else if b.isFourOfAKind() {
			value = FourOfAKind
		} else if b.isFullHouse() {
			value = FullHouse
		} else if b.isRoyalStraight() {
			value = Straight
		} else if b.isStraight() {
			value = Straight
		} else if b.isThreeOfAKind() {
			value = ThreeOfAKind
		}
		return
	case TwoPair: // 两对
		if b.isFlush() {
			value = Flush
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		} else if b.isFourOfAKind() {
			value = FourOfAKind
		} else if b.isFullHouse() {
			value = FullHouse
		} else if b.isRoyalStraight() {
			value = Straight
		} else if b.isStraight() {
			value = Straight
		} else if b.isThreeOfAKind() {
			value = ThreeOfAKind
		} else if b.isTwoPair() {
			value = TwoPair
		}
		return
	case OnePair: // 对子
		if b.isFlush() {
			value = Flush
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		} else if b.isFourOfAKind() {
			value = FourOfAKind
		} else if b.isFullHouse() {
			value = FullHouse
		} else if b.isRoyalStraight() {
			value = Straight
		} else if b.isStraight() {
			value = Straight
		} else if b.isThreeOfAKind() {
			value = ThreeOfAKind
		} else if b.isTwoPair() {
			value = TwoPair
		} else if b.isOnePair() {
			value = OnePair
		}
		return
	default:
		value = NoPair // 高牌
		if b.isFlush() {
			value = Flush
			if b.isRoyalStraight() {
				value = RoyalFlush
			} else if b.isStraight() {
				value = StraightFlush
			}
		} else if b.isFourOfAKind() {
			value = FourOfAKind
		} else if b.isFullHouse() {
			value = FullHouse
		} else if b.isRoyalStraight() {
			value = Straight
		} else if b.isStraight() {
			value = Straight
		} else if b.isThreeOfAKind() {
			value = ThreeOfAKind
		} else if b.isTwoPair() {
			value = TwoPair
		} else if b.isOnePair() {
			value = OnePair
		}
	}
	return
}

// 有序手牌牌值逐张比较(高牌，同花)
// result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (b *BoradCards) compareValue(a *BoradCards) int8 {
	if b[0].Value == a[0].Value &&
		b[1].Value == a[1].Value &&
		b[2].Value == a[2].Value &&
		b[3].Value == a[3].Value &&
		b[4].Value == a[4].Value {
		return 0
	}
	for i := 4; i >= 0; i-- {
		if a[i].Value > b[i].Value {
			return 1
		}
	}
	return -1
}

// 有序顺子比较
// result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (b *BoradCards) compareStraight(a *BoradCards) int8 {
	if b[0].Value == a[0].Value &&
		b[1].Value == a[1].Value &&
		b[2].Value == a[2].Value &&
		b[3].Value == a[3].Value &&
		b[4].Value == a[4].Value {
		return 0
	}
	if b[0].Value == 0x02 && b[4].Value == 0x0e {
		return 1
	}
	if a[0].Value == 0x02 && a[4].Value == 0x0e {
		return -1
	}
	if b[4].Value > a[4].Value {
		return -1
	}
	return 1
}

func (b *BoradCards) valueFourOfAKind() (int, int) {
	// 1111x
	if b[1].Value == b[0].Value &&
		b[2].Value == b[0].Value &&
		b[3].Value == b[0].Value &&
		b[4].Value != b[0].Value {
		return 0, 4
	}
	// x1111
	if b[0].Value != b[1].Value &&
		b[2].Value == b[1].Value &&
		b[3].Value == b[1].Value &&
		b[4].Value == b[1].Value {
		return 1, 0
	}
	return -1, -1
}

// 比较两个金刚大小
// result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (b *BoradCards) compareFourOfAKind(a *BoradCards) int8 {
	b1, b2 := b.valueFourOfAKind()
	a1, a2 := a.valueFourOfAKind()
	if b1 == -1 || a1 == -1 {
		return 0
	}
	if b[b1].Value > a[a1].Value {
		return -1
	}
	if b[b1].Value < a[a1].Value {
		return 1
	}
	if b[b2].Value == a[a2].Value {
		return 0
	}
	if b[b2].Value > a[a2].Value {
		return -1
	}
	return 1
}

func (b *BoradCards) valueFullHouse() (int, int) {
	// 111xx
	if b[0].Value == b[1].Value &&
		b[0].Value == b[2].Value &&
		b[0].Value != b[3].Value &&
		b[3].Value == b[4].Value {
		return 0, 3
	}
	// xx111
	if b[0].Value == b[1].Value &&
		b[0].Value != b[2].Value &&
		b[2].Value == b[3].Value &&
		b[2].Value == b[4].Value {
		return 2, 0
	}
	return -1, -1
}

// 比较两个葫芦大小
// result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (b *BoradCards) compareFullHouse(a *BoradCards) int8 {
	b1, b2 := b.valueFullHouse()
	a1, a2 := a.valueFullHouse()
	if a1 == -1 || b1 == -1 {
		return 0
	}
	if b[b1].Value > a[a1].Value {
		return -1
	}
	if b[b1].Value < a[a1].Value {
		return 1
	}
	if b[b2].Value == a[a2].Value {
		return 0
	}
	if b[b2].Value > a[a2].Value {
		return -1
	}
	return 1
}

func (b *BoradCards) valueThreeOfAKind() (int, int, int) {
	// 111xx
	if b[0].Value == b[1].Value &&
		b[0].Value == b[2].Value &&
		b[0].Value != b[3].Value &&
		b[0].Value != b[4].Value &&
		b[3].Value != b[4].Value {
		return 0, 4, 3
	}
	// x111x
	if b[0].Value != b[1].Value &&
		b[1].Value == b[2].Value &&
		b[1].Value == b[3].Value &&
		b[1].Value != b[4].Value &&
		b[0].Value != b[4].Value {
		return 1, 4, 0
	}
	// xx111
	if b[0].Value != b[2].Value &&
		b[1].Value != b[2].Value &&
		b[0].Value != b[1].Value &&
		b[2].Value == b[3].Value &&
		b[2].Value == b[4].Value {
		return 2, 1, 0
	}
	return -1, -1, -1
}

// 比较两个三条大小
// result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (b *BoradCards) compareThreeOfAKind(a *BoradCards) int8 {
	b1, b2, b3 := b.valueThreeOfAKind()
	a1, a2, a3 := a.valueThreeOfAKind()
	if a1 == -1 || b1 == -1 {
		return 0
	}
	if b[b1].Value > a[a1].Value {
		return -1
	}
	if b[b1].Value < a[a1].Value {
		return 1
	}
	if b[b2].Value == a[a2].Value && b[b3].Value == a[a3].Value {
		return 0
	}
	if b[b2].Value > a[a2].Value {
		return -1
	}
	if b[b2].Value < a[a2].Value {
		return 1
	}
	if b[b3].Value > a[a3].Value {
		return -1
	}
	return 1
}

func (b *BoradCards) valueTwoPair() (int, int, int) {
	// 1122x
	if b[0].Value == b[1].Value &&
		b[2].Value == b[3].Value &&
		b[0].Value != b[2].Value &&
		b[0].Value != b[4].Value &&
		b[2].Value != b[4].Value {
		return 2, 0, 4
	}
	// 11x33
	if b[0].Value == b[1].Value &&
		b[4].Value == b[3].Value &&
		b[0].Value != b[4].Value &&
		b[0].Value != b[2].Value &&
		b[2].Value != b[4].Value {
		return 3, 0, 2
	}
	// x2233
	if b[1].Value == b[2].Value &&
		b[4].Value == b[3].Value &&
		b[0].Value != b[4].Value &&
		b[0].Value != b[2].Value &&
		b[2].Value != b[4].Value {
		return 3, 1, 0
	}
	return -1, -1, -1
}

// 比较两个两对大小
// result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (b *BoradCards) compareTwoPair(a *BoradCards) int8 {
	b1, b2, b3 := b.valueTwoPair()
	a1, a2, a3 := a.valueTwoPair()
	if a1 == -1 || b1 == -1 {
		return 0
	}
	if b[b1].Value > a[a1].Value {
		return -1
	}
	if b[b1].Value < a[a1].Value {
		return 1
	}
	if b[b2].Value > a[a2].Value {
		return -1
	}
	if b[b2].Value < a[a2].Value {
		return 1
	}
	if b[b3].Value == a[a3].Value {
		return 0
	}
	if b[b3].Value > a[a3].Value {
		return -1
	}
	return 1
}

func (b *BoradCards) valueOnePair() (int, int, int, int) {
	// 11xxx
	if b[0].Value == b[1].Value &&
		b[0].Value != b[2].Value &&
		b[0].Value != b[3].Value &&
		b[0].Value != b[4].Value &&
		b[2].Value != b[3].Value &&
		b[2].Value != b[4].Value &&
		b[3].Value != b[4].Value {
		return 0, 4, 3, 2
	}
	// x11xx
	if b[1].Value == b[2].Value &&
		b[1].Value != b[0].Value &&
		b[1].Value != b[3].Value &&
		b[1].Value != b[4].Value &&
		b[0].Value != b[3].Value &&
		b[0].Value != b[4].Value &&
		b[3].Value != b[4].Value {
		return 1, 4, 3, 0
	}
	// xx11x
	if b[2].Value == b[3].Value &&
		b[2].Value != b[0].Value &&
		b[2].Value != b[1].Value &&
		b[2].Value != b[4].Value &&
		b[0].Value != b[1].Value &&
		b[0].Value != b[4].Value &&
		b[1].Value != b[4].Value {
		return 2, 4, 1, 0
	}
	// xxx11
	if b[3].Value == b[4].Value &&
		b[3].Value != b[0].Value &&
		b[3].Value != b[1].Value &&
		b[3].Value != b[2].Value &&
		b[0].Value != b[1].Value &&
		b[0].Value != b[2].Value &&
		b[1].Value != b[2].Value {
		return 3, 2, 1, 0
	}
	return -1, -1, -1, -1
}

// 比较两个两对大小
// result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (b *BoradCards) compareOnePair(a *BoradCards) int8 {
	b1, b2, b3, b4 := b.valueOnePair()
	a1, a2, a3, a4 := a.valueOnePair()
	if a1 == -1 || b1 == -1 {
		return 0
	}
	if b[b1].Value > a[a1].Value {
		return -1
	}
	if b[b1].Value < a[a1].Value {
		return 1
	}
	if b[b2].Value > a[a2].Value {
		return -1
	}
	if b[b2].Value < a[a2].Value {
		return 1
	}
	if b[b3].Value > a[a3].Value {
		return -1
	}
	if b[b3].Value < a[a3].Value {
		return 1
	}
	if b[b4].Value > a[a4].Value {
		return -1
	}
	return 1
}

// 比较牌力相同手牌大小
// result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func (b *BoradCards) compare(a *BoradCards, v int8) int8 {
	switch v {
	case NoPair: // 高牌
		return b.compareValue(a)
	case RoyalFlush: // 皇家同花顺
		return 1 // 不可能同时存在两个(即使是多人手牌)
	case StraightFlush: // 同花顺
		return b.compareStraight(a)
	case FourOfAKind: // 四条（金刚）
		return b.compareFourOfAKind(a)
	case FullHouse: // 葫芦
		return b.compareFullHouse(a)
	case Flush: // 同花
		return b.compareValue(a)
	case Straight: // 顺子
		return b.compareStraight(a)
	case ThreeOfAKind: // 三条
		return b.compareThreeOfAKind(a)
	case TwoPair: // 两对
		return b.compareTwoPair(a)
	case OnePair: // 对子
		return b.compareOnePair(a)
	}
	return -1
}

// 计算牌面最佳手牌及牌力
type BestCard struct {
	Value      int8       // 牌力大小
	Cards      BoradCards // 最佳手牌
	BoradCards BoradCards // 台面上公用牌
	HoleCards  HoleCards  // 手牌
}

// 摊牌计算牌力,并选择最佳手牌
func (b *BestCard) Showdown() {
	b7 := borad7Pool.Get().(*[7]Card)
	defer borad7Pool.Put(b7)
	b7[0], b7[1] = b.HoleCards[0], b.HoleCards[1]
	b7[2], b7[3], b7[4], b7[5], b7[6] = b.BoradCards[0], b.BoradCards[1], b.BoradCards[2], b.BoradCards[3], b.BoradCards[4]

	// 牌值升序排列(忽略花色)
	sort.Slice((*b7)[:], func(i, j int) bool {
		return (*b7)[i].Value < (*b7)[j].Value
	})

	b5 := borad5Pool.Get().(*BoradCards)
	defer borad5Pool.Put(b5)
	// 全部组合
	for _, i := range _riverCombinations {
		// 计算当前组合手牌牌力，并更新最佳手牌
		b5[0], b5[1], b5[2], b5[3], b5[4] = b7[i[0]], b7[i[1]], b7[i[2]], b7[i[3]], b7[i[4]]
		value := b5.getValue(b.Value)
		if value > b.Value {
			// 最佳手牌
			b.Value = value
			b.Cards = *b5
		} else if value == b.Value {
			// 选择最佳手牌
			com := b.Cards.compare(b5, value)
			if com > 0 {
				b.Cards = *b5
			}
		}
	}
}

// NewBestCard creates a new BestCard.
func NewBestCard(holeCards, boradCards []int16) *BestCard {
	b := &BestCard{}
	b.HoleCards[0], b.HoleCards[1] = _setCARDS[holeCards[0]], _setCARDS[holeCards[1]]
	b.BoradCards[0], b.BoradCards[1], b.BoradCards[2], b.BoradCards[3], b.BoradCards[4] = _setCARDS[boradCards[0]], _setCARDS[boradCards[1]], _setCARDS[boradCards[2]], _setCARDS[boradCards[3]], _setCARDS[boradCards[4]]
	return b
}
