func isHappy(n int) bool {
	for n != 1 && n != 4 {
		sum := 0
		for n > 0 {
			d := n % 10
			sum += d * d
			n = n / 10
		}
		n = sum
	}
	return n == 1
}
