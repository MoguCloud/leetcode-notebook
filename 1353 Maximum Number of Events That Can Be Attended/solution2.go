// 贪心 最小堆
//
// Your runtime beats 76.92 % of golang submissions (232 ms)
// Your memory usage beats 100 % of golang submissions (16.8 MB)
//
// 优先选择结束时间最近的会议。使用数组或者哈希表记录某天是否参加会议会超时，所以使用最小堆进行优化。

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	x := (*h)[:len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maxEvents(events [][]int) int {
	// events 按开始时间排序
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	days := 0
	for i := 0; i < len(events); i++ {
		if events[i][1] > days {
			days = events[i][1]
		}
	}
	h := &minHeap{}
	heap.Init(h)
	idx := 0
	count := 0
	for i := 1; i <= days; i++ {
		// 将从今天开始的event加入小顶堆中
		// 因为是将event[i][1]加入的小顶堆，所以堆顶元素为结束时间最小的event
		for idx < len(events) && events[idx][0] == i {
			heap.Push(h, events[idx][1])
			idx += 1
		}

		for len(*h) > 0 {
			// 堆顶元素也就是最小的结束日期<当前日期，从堆中移除；堆顶元素满足条件则从堆顶移除，并计数+1
			if (*h)[0] < i {
				heap.Pop(h)
			} else {
				heap.Pop(h)
				count += 1
				break
			}
		}
	}
	return count
}
