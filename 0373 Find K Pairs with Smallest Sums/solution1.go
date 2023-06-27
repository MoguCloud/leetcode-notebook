// Your runtime beats 25 % of golang submissions (268 ms)
// Your memory usage beats 63.64 % of golang submissions (10.4 MB)
//
// 小顶堆
//
// 时间复杂度 O(klogk)
// 空间复杂度 O(n)

type Pair struct {
	idx1  int
	idx2  int
	slice []int
}

type PairHeap []Pair

func (p PairHeap) Less(i, j int) bool {
	return p[i].slice[0]+p[i].slice[1] < p[j].slice[0]+p[j].slice[1]
}

func (p PairHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairHeap) Len() int {
	return len(p)
}

func (p *PairHeap) Push(x any) {
	*p = append(*p, x.(Pair))
}

func (p *PairHeap) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	set := make(map[string]struct{})
	ret := [][]int{}
	set[fmt.Sprintf("%d %d", 0, 0)] = struct{}{}
	h := &PairHeap{Pair{0, 0, []int{nums1[0], nums2[0]}}}
	heap.Init(h)
	for i := 0; i < k && i < len(nums1)*len(nums2); i++ {
		p := heap.Pop(h).(Pair)
		ret = append(ret, p.slice)
		if p.idx1+1 < len(nums1) {
			if _, ok := set[fmt.Sprintf("%d %d", p.idx1+1, p.idx2)]; !ok {
				set[fmt.Sprintf("%d %d", p.idx1+1, p.idx2)] = struct{}{}
				heap.Push(h, Pair{p.idx1 + 1, p.idx2, []int{nums1[p.idx1+1], nums2[p.idx2]}})
			}
		}
		if p.idx2+1 < len(nums2) {
			if _, ok := set[fmt.Sprintf("%d %d", p.idx1, p.idx2+1)]; !ok {
				set[fmt.Sprintf("%d %d", p.idx1, p.idx2+1)] = struct{}{}
				heap.Push(h, Pair{p.idx1, p.idx2 + 1, []int{nums1[p.idx1], nums2[p.idx2+1]}})
			}
		}
	}
	return ret
}
