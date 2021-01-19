// Time Limit Exceeded
//
// 大顶堆 将所有的球都放入堆内，取出堆顶，total += 球数，球数-1

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maxProfit(inventory []int, orders int) int {
	threshold := int(math.Pow(10, 9)) + 7
	h := &maxHeap{}
	heap.Init(h)
	for i := 0; i < len(inventory); i++ {
		heap.Push(h, inventory[i])
	}
	var count int64
	for orders > 0 {
		top := heap.Pop(h).(int)
		if len(*h) == 0 {
			// 没有其他的球了
			count += sum(int64(top), int64(orders))
			orders = 0
		} else {
			newTop := (*h)[0]
			delta := top - newTop + 1
			if delta > orders {
				delta = orders
			}
			count += sum(int64(top), int64(delta))
			orders -= delta
			heap.Push(h, top-delta)
		}
	}
	if count > int64(threshold) {
		count = count % int64(threshold)
	}
	return int(count)
}

func sum(num int64, i int64) int64 {
	return (num + num - i + 1) * i / 2
}
