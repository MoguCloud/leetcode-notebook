// Your runtime beats 17.73 % of golang submissions (334 ms)
// Your memory usage beats 48.20 % of golang submissions (10.6 MB)
//
// 堆
// 用一个大顶堆记录滑动窗口中的元素与索引
// 当滑动窗口移动时，堆顶元素的索引小于滑动窗口的右断点时，说明该元素已经不在滑动窗口中，需要从堆中移除
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

type MaxHeap [][2]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func maxSlidingWindow(nums []int, k int) []int {
	ret := []int{}
	pairs := [][2]int{}
	for i := 0; i < k-1; i++ {
		pairs = append(pairs, [2]int{nums[i], i})
	}
	hh := MaxHeap(pairs)
	h := &hh
	heap.Init(h)
	for i := k - 1; i < len(nums); i++ {
		heap.Push(h, [2]int{nums[i], i})
		for {
			x := (*h)[0]
			if x[1] > i-k {
				ret = append(ret, x[0])
				break
			}
			heap.Pop(h)
		}
	}
	return ret
}
