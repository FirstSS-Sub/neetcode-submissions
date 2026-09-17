func climbStairs(n int) int {
	memo := make([]int, n+1)
	return dfs(n, memo)
}

func dfs(n int, memo []int) int {
	if n <= 1 {
		return 1
	}
	
	if memo[n] != 0 {
		return memo[n]
	}
	
	memo[n] = dfs(n-1, memo) + dfs(n-2, memo)
	return memo[n]
}
