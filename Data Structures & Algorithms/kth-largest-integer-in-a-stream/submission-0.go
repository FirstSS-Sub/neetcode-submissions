type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k    int
	nums IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	ih := IntHeap(nums)
	heap.Init(&ih)

	for ih.Len() > k {
		heap.Pop(&ih)
	}

	return KthLargest{k: k, nums: ih}
}

func (this *KthLargest) Add(val int) int {
	if this.nums.Len() < this.k {
		heap.Push(&this.nums, val)
	} else if val > this.nums[0] {
		heap.Pop(&this.nums)
		heap.Push(&this.nums, val)
	}

	return this.nums[0]
}
