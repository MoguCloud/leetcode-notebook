// Your runtime beats 63.1 % of golang submissions (170 ms)
// Your memory usage beats 31.51 % of golang submissions (10.6 MB)
//
// 双指针
// 值交换
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func swapNodes(head *ListNode, k int) *ListNode {
	node := head
	for i := 1; i < k; i++ {
		node = node.Next
	}
	begin := node
	end := head
	node = node.Next
	for node != nil {
		node = node.Next
		end = end.Next
	}
	if end == begin {
		return head
	}
	begin.Val, end.Val = end.Val, begin.Val
	return head
}
