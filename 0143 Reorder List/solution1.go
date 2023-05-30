// Your runtime beats 71.38 % of golang submissions (8 ms)
// Your memory usage beats 88.75 % of golang submissions (5.3 MB)
//
// 翻转链表 合并链表
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func reorderList(head *ListNode) {
	// 通过快慢指针找到链表中点
	dummy := &ListNode{Next: head}
	s := dummy
	f := dummy
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	newNode := s.Next
	s.Next = nil
	// 翻转后半段链表
	var prev *ListNode
	for newNode != nil {
		tmp := newNode
		newNode = newNode.Next
		tmp.Next = prev
		prev = tmp
	}
	// 合并链表
	node := head
	n := prev
	for n != nil {
		newNode := node.Next
		node.Next = n
		n = n.Next
		node.Next.Next = newNode
		node = newNode
	}
}
