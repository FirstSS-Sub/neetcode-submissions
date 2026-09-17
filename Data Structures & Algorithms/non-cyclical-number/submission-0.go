func isHappy(n int) bool {
    set := make(map[int]struct{})
	set[n] = struct{}{}
	for {
		digits := getDigits(n)
		sum := 0
		for _, d := range digits {
			sum += d * d
		}
		if sum == 1 {
			return true
		}
		n = sum
		if _, ok := set[n]; ok {
			return false
		}
		set[n] = struct{}{}
	}
}

func getDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var digits []int
	for n > 0 {
		remainder := n % 10
		digits = append(digits, remainder)
		n = n / 10
	}

	return digits
}