import (
	"cmp"
	"slices"
)

type Item struct {
	num int
	count int
}

func topKFrequent(nums []int, k int) []int {
	var arr [2001]Item
	for _, n := range nums {
		arr[n + 1000].num = n
		arr[n + 1000].count += 1
	}
	slices.SortFunc(arr[:], func(a, b Item) int {
		return cmp.Compare(b.count, a.count)
	})
	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		item := arr[i]
		ans = append(ans, item.num)
	}
	return ans
}
