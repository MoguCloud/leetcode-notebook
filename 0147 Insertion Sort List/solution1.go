// Your runtime beats 98.55 % of golang submissions (2 ms)
// Your memory usage beats 40.58 % of golang submissions (3.3 MB)
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := dummyHead.Next
	for prev.Next != nil {
		if prev.Next.Val >= prev.Val {
			prev = prev.Next
		} else {
			for node := dummyHead; node.Next != prev.Next; node = node.Next {
				if prev.Next.Val < node.Next.Val {
					curr := prev.Next
					prev.Next = curr.Next
					curr.Next = node.Next
					node.Next = curr
					break
				}
			}
		}
	}
	return dummyHead.Next
}
