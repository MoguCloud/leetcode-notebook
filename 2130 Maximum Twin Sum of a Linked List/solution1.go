// Your runtime beats 48.5 % of golang submissions (202 ms)
// Your memory usage beats 85.6 % of golang submissions (9.4 MB)
//
// 使用数组存孪生和，求出最大值
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func pairSum(head *ListNode) int {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	sum := make([]int, n/2)
	node := head
	for i := 0; i <= (n/2)-1; i++ {
		sum[i] = node.Val
		node = node.Next
	}
	max := 0
	for i := (n / 2) - 1; i >= 0; i-- {
		s := sum[i] + node.Val
		if s > max {
			max = s
		}
		node = node.Next
	}
	return max
}
