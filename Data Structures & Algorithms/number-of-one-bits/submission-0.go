func hammingWeight(n int) int {
	div := 1 << 31
	bit := 0
	for div > 0 {
		if n / div == 1 {
			n -= div
			bit += 1
		}
		div >>= 1
	}
	return bit
}
