// Your runtime beats 62.89 % of golang submissions (2 ms)
// Your memory usage beats 85.57 % of golang submissions (2.1 MB)
//
// 迭代
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	i := dummy
	j := dummy
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		i = prev.Next
		j = prev.Next.Next
		i.Next = j.Next
		j.Next = i
		prev.Next = j
		prev = prev.Next.Next
	}
	return dummy.Next
}
