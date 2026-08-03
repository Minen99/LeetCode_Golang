package coinChange

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -2 // means not calculate
	}

	var solve func(rem int) int
	solve = func(rem int) int {
		// base case
		if rem == 0 {
			return 0
		}

		// already calculated
		if memo[rem] != -2 {
			return memo[rem]
		}

		// lets calculate
		best := -1
		for _, coin := range coins {
			if rem-coin < 0 { // is it possible coin?
				continue
			}

			sub := solve(rem - coin) // DFS
			if sub == -1 {           // not found
				continue
			}
			if best == -1 {
				best = sub + 1
			} else {
				best = min(best, sub+1)
			}
		}
		memo[rem] = best // memorize the result
		return best
	}

	return solve(amount)
}
