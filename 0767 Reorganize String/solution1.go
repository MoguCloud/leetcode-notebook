// Your runtime beats 72.55 % of golang submissions (1 ms)
// Your memory usage beats 70.59 % of golang submissions (2 MB)
//
// 大顶堆
// 当一个字母出现的次数大于 (len(s) + 1 ) / 2 时，则不满足条件
// 用一个堆维护字母出现的次数，堆顶为出现次数最多的字母
// 每次将剩余数量最多的字母插入新字符串中，如果堆顶元素和新字符串最后一个字母一样，则选择堆中第二大的
// 每次插入到新字符串后需要将字母出现的次数减去1再放回堆中，保证堆顶是剩余次数最多的
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

type charHeap []int

func (h charHeap) Len() int           { return len(h) }
func (h charHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h charHeap) Less(i, j int) bool { return charCount[h[i]] > charCount[h[j]] }
func (h *charHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h *charHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *charHeap) Top() int {
	return (*h)[0]
}

var charCount [26]int

func reorganizeString(s string) string {
	charCount = [26]int{}
	for i := 0; i < len(s); i++ {
		charCount[s[i]-'a']++
	}
	for i := 0; i < len(charCount); i++ {
		if charCount[i] > (len(s)+1)/2 {
			return ""
		}
	}
	intSlice := make([]int, 26)
	for i := 0; i < 26; i++ {
		intSlice[i] = i
	}
	hh := charHeap(intSlice)
	h := &hh
	heap.Init(h)
	ret := make([]byte, len(s))
	prev := -1
	for i := 0; i < len(ret); i++ {
		var next int
		if prev == h.Top() {
			heap.Pop(h)
			next = heap.Pop(h).(int)
			heap.Push(h, prev)
		} else {
			next = heap.Pop(h).(int)
		}
		ret[i] = byte('a' + next)
		charCount[next]--
		heap.Push(h, next)
		prev = next
	}
	return string(ret)
}
