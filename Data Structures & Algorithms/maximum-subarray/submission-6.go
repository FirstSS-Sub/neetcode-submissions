func maxSubArray(nums []int) int {
    maxSub, curSum := nums[0], 0
    for _, num := range nums {
        if curSum < 0 {
            curSum = 0
        }
        curSum += num
        maxSub = max(maxSub, curSum)
    }
    return maxSub
}