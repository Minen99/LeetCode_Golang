package coinChange

func coinChange(coins []int, amount int) int {
	const INF = 1 << 30

	dp := make([]int, amount+1)
	// initialize the dp array
	for i := 1; i < (amount + 1); i++ {
		dp[i] = INF
	}

	// fill in dp in ascending order
	for x := 0; x < (amount + 1); x++ {
		for _, coin := range coins {
			if coin <= x && (dp[x-coin]+1) < dp[x] {
				dp[x] = dp[x-coin] + 1
			}
		}
	}

	// if there is no correct combinations
	if dp[amount] == INF {
		return -1
	} else {
		return dp[amount]
	}
}
