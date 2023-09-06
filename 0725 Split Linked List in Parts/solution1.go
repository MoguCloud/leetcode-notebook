// Your runtime beats 56.52 % of golang submissions (3 ms)
// Your memory usage beats 56.52 % of golang submissions (2.77 MB)
//
// 模拟
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func splitListToParts(head *ListNode, k int) []*ListNode {
	count := 0
	for node := head; node != nil; node = node.Next {
		count++
	}
	m := count / k
	n := count - m*k
	// n 个 长度为 (m + 1), (k - n) 个长度为 m 的
	ret := make([]*ListNode, k)
	node := head
	offset := 0
	var prev *ListNode
	for i := 0; i < n; i++ {
		ret[offset] = node
		offset++
		for j := 0; j < m+1; j++ {
			prev = node
			node = node.Next
		}
		prev.Next = nil
	}
	for i := 0; i < k-n; i++ {
		ret[offset] = node
		offset++
		for j := 0; j < m; j++ {
			prev = node
			node = node.Next
		}
		if prev != nil {
			prev.Next = nil
		}
	}
	return ret
}
