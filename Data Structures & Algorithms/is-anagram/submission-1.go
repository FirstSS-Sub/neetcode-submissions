func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int, len(s))
	for _, r := range s {
		m[r] += 1
	}
	for _, r := range t {
		m[r] -= 1
		if m[r] < 0 {
			return false
		}
	}
	return true
}
