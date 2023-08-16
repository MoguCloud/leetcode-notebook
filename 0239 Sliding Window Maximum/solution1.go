// Your runtime beats 14.68 % of golang submissions (396 ms)
// Your memory usage beats 48.20 % of golang submissions (10.6 MB)
//
// 堆
// 用一个大顶堆A记录滑动窗口中的元素
// 用一个大顶堆B记录滑出窗口的元素
// 当2个堆顶元素相同时，说明堆A的堆顶元素已经不在滑动窗口里了
// 需要从堆里移除，直到2个堆顶的元素不同
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func maxSlidingWindow(nums []int, k int) []int {
	ret := []int{}
	prevHeap := &MaxHeap{}
	heap.Init(prevHeap)
	slice := make([]int, k-1)
	copy(slice, nums[:k-1])
	ch := MaxHeap(slice)
	currHeap := &ch
	heap.Init(currHeap)
	for i := k - 1; i < len(nums); i++ {
		heap.Push(currHeap, nums[i])
		if i-k >= 0 {
			heap.Push(prevHeap, nums[i-k])
		}
		for prevHeap.Len() > 0 && (*prevHeap)[0] == (*currHeap)[0] {
			heap.Pop(prevHeap)
			heap.Pop(currHeap)
		}
		ret = append(ret, (*currHeap)[0])
	}
	return ret
}
