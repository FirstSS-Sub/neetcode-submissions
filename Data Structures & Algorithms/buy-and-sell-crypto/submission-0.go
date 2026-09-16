func maxProfit(prices []int) int {
	ans := 0
	buy := 101
	for _, n := range prices {
		ans = max(ans, n-buy)
		buy = min(buy, n)
	}
	return ans
}
