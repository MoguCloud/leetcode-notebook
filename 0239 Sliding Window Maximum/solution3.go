// Your runtime beats 83.10 % of golang submissions (222 ms)
// Your memory usage beats 38.23 % of golang submissions (11.70 MB)
//
// 单调队列
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

type MonotonicQueue []int

func (q MonotonicQueue) max() int {
	return q[0]
}

func (q *MonotonicQueue) push(x int) {
	for len(*q) > 0 && (*q)[len(*q)-1] < x {
		*q = (*q)[:len(*q)-1]
	}
	*q = append((*q), x)
}

func (q *MonotonicQueue) pop(x int) {
	if len(*q) > 0 && (*q)[0] == x {
		*q = (*q)[1:len(*q)]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	ret := []int{}
	q := &MonotonicQueue{}
	for i := 0; i < k-1; i++ {
		q.push(nums[i])
	}
	for i := k - 1; i < len(nums); i++ {
		q.push(nums[i])
		if i >= k {
			q.pop(nums[i-k])
		}
		ret = append(ret, q.max())
	}
	return ret
}
