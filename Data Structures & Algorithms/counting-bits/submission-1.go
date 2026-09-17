func countBits(n int) []int {
	list := make([]int, n+1)
	for i := 0; i <= n; i++ {
		div := 1 << 10
		bit := 0
		ii := i
		for div > 0 {
			if ii / div == 1 {
				ii -= div
				bit += 1
			}
			div >>= 1
		}
		list[i] = bit
	}
	return list
}
