func maxSubArray(nums []int) int {
	cm := make([]int, len(nums))
	sum := 0
	mn, mx, numsMax := 1000000001, -1000000001, -1000000001
	mnI, mxI := 0, 0
	for i, n := range nums {
		numsMax = max(numsMax, n)
		sum += n
		cm[i] = sum
		if mn > sum {
			mn = sum
			mnI = i
		}
		if mx <= sum {
			mx = sum
			mxI = i
		}
	}
	if numsMax < 0 {
		return numsMax
	}
	if mn < 0 && mnI < mxI {
		return mx - mn
	}
	mn2 := 0
	for i := 0; i < mxI; i++ {
		mn2 = min(mn2, cm[i])
	}
	return mx - mn2
}