// Last updated: 6/27/2025, 4:40:27 PM
type Item struct {
	value    int
	priority int
}

type MinHeap []*Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	item := x.(*Item)
	*h = append(*h, item)
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	minHeap := &MinHeap{}
	for value, count := range frequency {
		heap.Push(minHeap, &Item{value: value, priority: count})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	result := make([]int, 0, k)
	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(*Item)
		result = append(result, item.value)
	}

	return result
}
