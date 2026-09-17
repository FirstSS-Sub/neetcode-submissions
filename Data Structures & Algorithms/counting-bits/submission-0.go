func countBits(n int) []int {
	list := make([]int, 0, n)
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
		list = append(list, bit)
	}
	return list
}
