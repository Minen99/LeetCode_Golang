package coinChange

import "sort"

type coinCnt struct {
	kind   int
	pieces int
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	originalAmount := amount
	coinCounts := []coinCnt{}

	sort.Slice(coins, func(i int, j int) bool {
		return coins[j] < coins[i]
	})
	for _, coin := range coins {
		count := coinCnt{kind: coin}
		count.pieces = amount / coin
		amount = amount % coin
		coinCounts = append(coinCounts, count)
	}

	sum := 0
	numOfCoins := 0
	for _, cnt := range coinCounts {
		sum += cnt.kind * cnt.pieces
		numOfCoins += cnt.pieces
	}

	if originalAmount == sum {
		return numOfCoins
	} else {
		return -1
	}
}
