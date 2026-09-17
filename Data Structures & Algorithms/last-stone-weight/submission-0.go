type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	maxHeap := MaxHeap(stones)
	heap.Init(&maxHeap)
	for maxHeap.Len() > 1 {
		big := heap.Pop(&maxHeap).(int)
		small := heap.Pop(&maxHeap).(int)
		result := big - small
		if result != 0 {
			heap.Push(&maxHeap, result)
		}
	}
	if maxHeap.Len() == 1 {
		return maxHeap.Pop().(int)
	}
	return 0
}
