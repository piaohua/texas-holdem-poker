package texasholdempoker

import (
	"testing"
)

func Test_isFlush(t *testing.T) {
	var l = []struct {
		cards  BoradCards
		result bool
	}{
		{
			BoradCards{
				{SPADE, 2},
				{SPADE, 3},
				{SPADE, 4},
				{SPADE, 5},
				{SPADE, 6},
			},
			true,
		},
		{
			BoradCards{
				{SPADE, 2},
				{DIAMOND, 3},
				{SPADE, 4},
				{SPADE, 5},
				{SPADE, 6},
			},
			false,
		},
	}
	for _, v := range l {
		if v.cards.isFlush() != v.result {
			t.Fatal(v.cards)
		}
	}
}

func Test_isFlush2(t *testing.T) {
	var l = []struct {
		cards  [5]int16
		result bool
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, true},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, true},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, true},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, false},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, false},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, false},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, false},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, true},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, true},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, true},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, true},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, false},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, false},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, false},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, false},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, false},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, false},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, false},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.isFlush() != v.result {
			t.Fatal(b)
		}
	}
}

func Test_isRoyalStraight(t *testing.T) {
	var l = []struct {
		cards  [5]int16
		result bool
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, true},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, false},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, false},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, false},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, false},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, false},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, false},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, false},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, false},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, false},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, false},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, false},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, false},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, false},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, false},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, false},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.isRoyalStraight() != v.result {
			t.Fatal(b)
		}
	}
}

func Test_isFourOfAKind(t *testing.T) {
	var l = []struct {
		cards  [5]int16
		result bool
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, false},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, false},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, true},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, true},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, false},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, false},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, false},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, false},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, false},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, false},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, false},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, false},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, false},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, false},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, false},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, false},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.isFourOfAKind() != v.result {
			t.Fatal(b)
		}
	}
}

func Test_isFullHouse(t *testing.T) {
	var l = []struct {
		cards  [5]int16
		result bool
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, false},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, false},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, false},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, false},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, true},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, true},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, false},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, false},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, false},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, false},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, false},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, false},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, false},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, false},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, false},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, false},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.isFullHouse() != v.result {
			t.Fatal(b)
		}
	}
}

func Test_isStraight(t *testing.T) {
	var l = []struct {
		cards  [5]int16
		result bool
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, true},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, true},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, true},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, false},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, false},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, false},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, false},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, false},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, false},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, false},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, true},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, true},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, false},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, false},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, false},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, false},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, false},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, false},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.isStraight() != v.result {
			t.Fatal(b)
		}
	}
}

func Test_isThreeOfAKind(t *testing.T) {
	var l = []struct {
		cards  [5]int16
		result bool
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, false},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, false},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, false},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, false},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, false},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, false},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, false},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, false},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, false},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, false},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, true},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, true},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, true},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, false},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, false},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, false},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, false},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, false},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.isThreeOfAKind() != v.result {
			t.Fatal(b)
		}
	}
}

func Test_isTwoPair(t *testing.T) {
	var l = []struct {
		cards  [5]int16
		result bool
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, false},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, false},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, false},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, false},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, false},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, false},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, false},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, false},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, false},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, false},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, false},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, true},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, true},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, true},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, false},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, false},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, false},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.isTwoPair() != v.result {
			t.Fatal(b)
		}
	}
}

func Test_isOnePair(t *testing.T) {
	var l = []struct {
		cards  [5]int16
		result bool
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, false},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, false},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, false},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, false},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, false},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, false},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, false},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, false},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, false},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, false},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, false},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, false},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, false},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, false},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, false},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, false},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, false},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, true},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, true},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, true},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, true},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, false},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.isOnePair() != v.result {
			t.Fatal(b)
		}
	}
}

func Test_Showdown(t *testing.T) {
	b := &BestCard{
		HoleCards: HoleCards{
			{SPADE, 3},
			{HEART, 4},
		},
		BoradCards: BoradCards{
			{HEART, 5},
			{SPADE, 6},
			{CLUB, 7},
			{DIAMOND, 8},
			{HEART, 9},
		},
	}
	b.Showdown()
	t.Log(b.Cards.String())
}

func Test_getValue(t *testing.T) {
	var l = []struct {
		cards [5]int16
		value int8
	}{
		// 皇家同花顺
		{[5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, RoyalFlush},
		// 同花顺
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, StraightFlush},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, StraightFlush},
		// 4条
		{[5]int16{0x1a, 0x2a, 0x3a, 0x4a, 0x1e}, FourOfAKind},
		{[5]int16{0x1a, 0x2e, 0x3e, 0x4e, 0x1e}, FourOfAKind},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x32, 0x44, 0x14}, FullHouse},
		{[5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, FullHouse},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, Flush},
		{[5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, Flush},
		{[5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, Flush},
		{[5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, Flush},
		// 顺子
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, Straight},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, Straight},
		// 三条
		{[5]int16{0x12, 0x22, 0x32, 0x45, 0x1e}, ThreeOfAKind},
		{[5]int16{0x13, 0x24, 0x34, 0x44, 0x1e}, ThreeOfAKind},
		{[5]int16{0x13, 0x24, 0x35, 0x45, 0x15}, ThreeOfAKind},
		// 两对
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4e, 0x1e}, TwoPair},
		{[5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, TwoPair},
		{[5]int16{0x1a, 0x2b, 0x3b, 0x4e, 0x1e}, TwoPair},
		// 对子
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x16}, OnePair},
		{[5]int16{0x12, 0x23, 0x33, 0x44, 0x15}, OnePair},
		{[5]int16{0x12, 0x23, 0x3b, 0x4b, 0x1e}, OnePair},
		{[5]int16{0x12, 0x23, 0x34, 0x4e, 0x1e}, OnePair},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, HighCard},
	}
	b := borad5Pool.Get().(*BoradCards)
	for _, v := range l {
		b[0], b[1], b[2], b[3], b[4] = _setCARDS[v.cards[0]], _setCARDS[v.cards[1]], _setCARDS[v.cards[2]], _setCARDS[v.cards[3]], _setCARDS[v.cards[4]]
		if b.getValue(0) != v.value {
			t.Fatal(b)
		}
	}
}

func Test_ShowdownCompare(t *testing.T) {
	var l = []struct {
		borad [5]int16
		hole  [2]int16
		cards [5]int16
		value int8
	}{
		// 皇家同花顺
		{[5]int16{0x12, 0x13, 0x1c, 0x1d, 0x1e}, [2]int16{0x1a, 0x1b}, [5]int16{0x1a, 0x1b, 0x1c, 0x1d, 0x1e}, RoyalFlush},
		// 同花顺
		{[5]int16{0x42, 0x13, 0x14, 0x15, 0x26}, [2]int16{0x12, 0x16}, [5]int16{0x12, 0x13, 0x14, 0x15, 0x16}, StraightFlush},
		{[5]int16{0x22, 0x23, 0x2a, 0x2c, 0x2e}, [2]int16{0x24, 0x25}, [5]int16{0x22, 0x23, 0x24, 0x25, 0x2e}, StraightFlush},
		// 4条
		{[5]int16{0x19, 0x2a, 0x3a, 0x4c, 0x1e}, [2]int16{0x1a, 0x4a}, [5]int16{0x1a, 0x4a, 0x2a, 0x3a, 0x1e}, FourOfAKind},
		{[5]int16{0x13, 0x24, 0x3a, 0x4e, 0x1e}, [2]int16{0x2e, 0x3e}, [5]int16{0x3a, 0x2e, 0x3e, 0x4e, 0x1e}, FourOfAKind},
		// 葫芦
		{[5]int16{0x12, 0x22, 0x33, 0x44, 0x19}, [2]int16{0x32, 0x14}, [5]int16{0x32, 0x12, 0x22, 0x14, 0x44}, FullHouse},
		{[5]int16{0x1a, 0x22, 0x39, 0x45, 0x15}, [2]int16{0x12, 0x35}, [5]int16{0x12, 0x22, 0x35, 0x45, 0x15}, FullHouse},
		// 同花
		{[5]int16{0x12, 0x13, 0x14, 0x35, 0x37}, [2]int16{0x15, 0x17}, [5]int16{0x12, 0x13, 0x14, 0x15, 0x17}, Flush},
		{[5]int16{0x22, 0x23, 0x24, 0x15, 0x37}, [2]int16{0x25, 0x27}, [5]int16{0x22, 0x23, 0x24, 0x25, 0x27}, Flush},
		{[5]int16{0x32, 0x33, 0x34, 0x15, 0x47}, [2]int16{0x35, 0x37}, [5]int16{0x32, 0x33, 0x34, 0x35, 0x37}, Flush},
		{[5]int16{0x42, 0x43, 0x44, 0x15, 0x37}, [2]int16{0x45, 0x47}, [5]int16{0x42, 0x43, 0x44, 0x45, 0x47}, Flush},
		// 顺子
		{[5]int16{0x23, 0x34, 0x45, 0x1a, 0x3e}, [2]int16{0x12, 0x16}, [5]int16{0x12, 0x23, 0x34, 0x45, 0x16}, Straight},
		{[5]int16{0x12, 0x23, 0x34, 0x49, 0x1a}, [2]int16{0x45, 0x1e}, [5]int16{0x12, 0x23, 0x34, 0x45, 0x1e}, Straight},
		// 三条
		{[5]int16{0x12, 0x22, 0x33, 0x45, 0x1e}, [2]int16{0x19, 0x32}, [5]int16{0x32, 0x12, 0x22, 0x19, 0x1e}, ThreeOfAKind},
		{[5]int16{0x12, 0x23, 0x25, 0x35, 0x1c}, [2]int16{0x14, 0x45}, [5]int16{0x14, 0x45, 0x25, 0x35, 0x1c}, ThreeOfAKind},
		{[5]int16{0x13, 0x24, 0x35, 0x48, 0x18}, [2]int16{0x22, 0x38}, [5]int16{0x24, 0x35, 0x38, 0x48, 0x18}, ThreeOfAKind},
		// 两对
		{[5]int16{0x17, 0x28, 0x3b, 0x4c, 0x1e}, [2]int16{0x27, 0x4e}, [5]int16{0x27, 0x17, 0x4c, 0x4e, 0x1e}, TwoPair},
		{[5]int16{0x13, 0x2a, 0x3b, 0x4b, 0x1e}, [2]int16{0x1a, 0x35}, [5]int16{0x1a, 0x2a, 0x3b, 0x4b, 0x1e}, TwoPair},
		{[5]int16{0x13, 0x25, 0x3b, 0x4c, 0x1e}, [2]int16{0x2b, 0x4e}, [5]int16{0x2b, 0x3b, 0x4c, 0x4e, 0x1e}, TwoPair},
		// 对子
		{[5]int16{0x13, 0x24, 0x35, 0x48, 0x1a}, [2]int16{0x12, 0x22}, [5]int16{0x12, 0x22, 0x35, 0x48, 0x1a}, OnePair},
		{[5]int16{0x12, 0x23, 0x34, 0x45, 0x19}, [2]int16{0x35, 0x17}, [5]int16{0x34, 0x35, 0x45, 0x17, 0x19}, OnePair},
		{[5]int16{0x12, 0x23, 0x35, 0x46, 0x1e}, [2]int16{0x3b, 0x4b}, [5]int16{0x35, 0x46, 0x3b, 0x4b, 0x1e}, OnePair},
		{[5]int16{0x13, 0x25, 0x37, 0x29, 0x4e}, [2]int16{0x12, 0x1e}, [5]int16{0x25, 0x37, 0x29, 0x1e, 0x4e}, OnePair},
		// 高牌
		{[5]int16{0x12, 0x23, 0x3b, 0x4d, 0x1e}, [2]int16{0x48, 0x47}, [5]int16{0x47, 0x48, 0x3b, 0x4d, 0x1e}, HighCard},
	}

	for _, v := range l {
		b := &BestCard{}
		b.HoleCards[0], b.HoleCards[1] = _setCARDS[v.hole[0]], _setCARDS[v.hole[1]]
		b.BoradCards[0], b.BoradCards[1], b.BoradCards[2], b.BoradCards[3], b.BoradCards[4] = _setCARDS[v.borad[0]], _setCARDS[v.borad[1]], _setCARDS[v.borad[2]], _setCARDS[v.borad[3]], _setCARDS[v.borad[4]]
		b.Showdown()
		if b.Value != v.value {
			t.Fatalf("invalid value: %d, v: %v", b.Value, v)
		}
		for i := range v.cards {
			if _setCARDS[v.cards[i]] != b.Cards[i] {
				t.Fatalf("invalid best cards: %s != %v, %s,%s, %d", b.Cards, v.cards, b.Cards[i], _setCARDS[v.cards[i]], i)
			}
		}
	}
}

func Test_NewBestCard(t *testing.T) {
	holeCards := [2]int16{0x13, 0x14}
	boradCards := [5]int16{0x13, 0x14, 0x15, 0x16, 0x17}
	b := NewBestCard(holeCards[:], boradCards[:])
	t.Log(b.HoleCards)
	t.Log(b.BoradCards)
}
