func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		now := (left+right) / 2
		if nums[now] == target {
			return now
		} else if nums[now] > target {
			right = now-1
		} else {
			left = now+1
		}
	}
	return -1
}
