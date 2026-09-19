type Solution struct{}

var hoge []string

func (s *Solution) Encode(strs []string) string {
	hoge = strs
	ans := ""
	for _, s := range strs {
		ans += s
	}
	return ans
}

func (s *Solution) Decode(encoded string) []string {
	return hoge
}
