// Your runtime beats 14.29 % of golang submissions (936 ms)
// Your memory usage beats 9.52 % of golang submissions (236.9 MB)
//
// 小顶堆
//
// 时间复杂度 O(m*n*log(m*n))
// 空间复杂度 O(m*n)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func nthSuperUglyNumber(n int, primes []int) int {
    if n == 1 {
        return 1
    }
    h := &IntHeap{}
    heap.Init(h)
    for _, p := range primes {
        heap.Push(h, p)
        for _, pp := range primes {
            heap.Push(h, p * pp)
        }
    }
    for n > 2 {
        p := heap.Pop(h).(int)
        for len(*h) > 0 && (*h)[0] == p {
            heap.Pop(h)
        }
        for _, pp := range primes {
            heap.Push(h, p * pp)
        }
        n--
    }
    p := heap.Pop(h).(int)
    return p
}
