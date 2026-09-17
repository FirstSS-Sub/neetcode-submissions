func minCostClimbingStairs(cost []int) int {
    memo := make([]int, len(cost)+1)
	return dp(len(cost), cost, memo)
}

func dp(i int, cost []int, memo []int) int {
	if i <= 1 {
		return 0
	}

	if memo[i] != 0 {
		return memo[i]
	}

	down1 := dp(i-1, cost, memo) + cost[i-1]
	down2 := dp(i-2, cost, memo) + cost[i-2]

	if down1 < down2 {
		memo[i] = down1
	} else {
		memo[i] = down2
	}
	return memo[i]
}