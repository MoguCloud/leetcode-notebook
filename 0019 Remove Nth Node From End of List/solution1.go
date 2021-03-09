// 快慢指针
//
// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 25.77 % of golang submissions (2.2 MB)
//
// 快慢指针都从头开始，先让快指针走 n 步，然后快慢指针同时前进。
// 当快指针为 null 的时候，慢指针停留在倒数第 n 个节点。此时删除慢指针所在的节点即可。
// slow.Prev.Next = slow.Next 因为没有 Prev 指针，所以在当快指针走到 null 时，慢指针不走，此时慢指针停留在 Prev 处。
// 此时将 slow.Next = slow.Next.Next 即可将节点删除

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for {
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
