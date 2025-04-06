## Texas Hold’em poker in Go

用Golang实现德州扑克游戏牌型，洗牌和最佳牌型计算。

### 定义实现

牌型花色定义
```go
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
```

牌值定义从`0x02 - 0x0e`, `0x0b - 0x0e`分别表示`J,Q,K,A`
```go
const (
    TWO int8 = iota + 2
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
    ACE
)
```
用`0x0e`表示`ACE`，可以方便后续的大小比较，然后最小顺子`(A,2,3,4,5)`特殊处理即可。

牌值使用`int16`，高位表示花色，低位表示牌值
```go
	// color: 0x4e >> 4   = 4
	// value: 0x4e & 0x0f = 0x0e
    // (0x04 << 4) | 0x0e = 0x4e
```

最终牌值计算和比较时，需要将`int16`表示转换为`Card`类型，方便处理。

### 测试

执行洗牌测试用例
```sh
go test -v -run="Test_Shuffle"
```

执行计算牌力和最佳手牌用例
```sh
go test -v -run="Test_Showdown"

go test -v -run="Test_ShowdownCompare"
```

### 参考
- [Shuffling a Deck of Cards in Go: Secure Randomness with crypto/rand][1]
- [德州扑克术语][2]

[1]: https://medium.com/@suri.podeti7/shuffling-a-deck-of-cards-in-go-secure-randomness-with-crypto-rand-fe5e1584645b
[2]: https://mbd.baidu.com/newspage/data/dtlandingsuper?nid=dt_5986696310088131183&sourceFrom=search_a