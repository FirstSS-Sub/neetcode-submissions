func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		m[n] = struct{}{}
	}
	ans := 0
	for _, n := range nums {
		if _, ok1 := m[n-1]; !ok1 {
			i := n
			ok := true
			for ok {
				i += 1
				_, ok = m[i]
			}
			ans = max(ans, i-n)
		}
	}
	return ans
}
