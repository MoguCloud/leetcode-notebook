// Your runtime beats 73.33 % of golang submissions (144 ms)
// Your memory usage beats 33.33 % of golang submissions (11.6 MB)
//
// 小顶堆
//
// 时间复杂度 O(klogn)
// 空间复杂度 O(candidates)

// 第0个元素用来判断是左侧加入的还是右侧加入的，第1个元素是雇佣工人的代价
type WorkerHeap [][2]int

func (h WorkerHeap) Len() int {
	return len(h)
}

func (h WorkerHeap) Less(i, j int) bool {
	if h[i][1] < h[j][1] {
		return true
	} else if h[i][1] == h[j][1] {
		return h[i][0] < h[j][0]
	}
	return false
}

func (h WorkerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WorkerHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *WorkerHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	h := &WorkerHeap{}
	var i, j int
	// 左侧元素加入堆中
	for i = 0; i < candidates; i++ {
		h.Push([2]int{0, costs[i]})
	}
	// 右侧元素加入堆中
	for j = len(costs) - 1; j > len(costs)-1-candidates && j >= i; j-- {
		h.Push([2]int{1, costs[j]})
	}
	heap.Init(h)
	var ret int64
	for k > 0 {
		w := heap.Pop(h).([2]int)
		ret += int64(w[1])
		k--
		// 判断是否所有元素都加入过堆中
		if i > j {
			continue
		}
		// 判断最小代价的是左侧还是右侧的，来决定后续从左侧还是右侧加入堆中
		if w[0] == 0 {
			heap.Push(h, [2]int{0, costs[i]})
			i++
		} else {
			heap.Push(h, [2]int{1, costs[j]})
			j--
		}
	}
	return ret
}
