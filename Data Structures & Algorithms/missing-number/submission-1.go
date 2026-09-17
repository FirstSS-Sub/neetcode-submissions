func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums)+ 1) / 2
	for _, n := range nums {
		sum -= n
	}
	return sum
}
