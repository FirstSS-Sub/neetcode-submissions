func productExceptSelf(nums []int) []int {
	product := 1
	zeroFlag := false
	zeroCount := 0
	for _, n := range nums {
		if n != 0 {
			product *= n
		} else {
			zeroFlag = true
			zeroCount += 1
		}
	}
	if len(nums) == zeroCount || zeroCount >= 2 {
		return make([]int, len(nums))
	}
	ans := make([]int, 0, len(nums))
	for _, n := range nums {
		if n != 0 && !zeroFlag {
			ans = append(ans, product / n)
		} else if n != 0 {
			ans = append(ans, 0)
		} else {
			ans = append(ans, product)
		}
	}
	return ans
}
