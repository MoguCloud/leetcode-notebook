// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (2.1 MB)

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Next: head}
	beforeM := dummy
	for i := 1; i < m; i++ {
		beforeM = beforeM.Next
	}
	pre := beforeM
	cur := beforeM.Next
	nodeN := beforeM.Next
	var nex *ListNode
	for i := m; i <= n; i++ {
		nex = cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	beforeM.Next = pre
	nodeN.Next = nex
	return dummy.Next
}
