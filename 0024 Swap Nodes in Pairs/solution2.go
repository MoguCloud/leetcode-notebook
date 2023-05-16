// Your runtime beats 62.89 % of golang submissions (2 ms)
// Your memory usage beats 85.57 % of golang submissions (2.1 MB)
//
// 递归
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
