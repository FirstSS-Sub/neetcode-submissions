func reverseBits(n int) int {
	div := 1 << 31
	i := 31
	res := 0
	for div > 0 {
		if n / div == 1 {
			n -= div
			res += (1 << (31-i))
		}
		div >>= 1
		i -= 1
	}
	return res
}
