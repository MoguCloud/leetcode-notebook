// Your runtime beats 59.26 % of golang submissions (11 ms)
// Your memory usage beats 70.37 % of golang submissions (4.6 MB)
//
// 翻转链表
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseListNode(l1)
	l2 = reverseListNode(l2)
	carry := 0
	dummy := &ListNode{}
	node := dummy
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = 0
		if sum >= 10 {
			carry = 1
		}
		newNode := &ListNode{Val: sum % 10}
		node.Next = newNode
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + carry
		carry = 0
		if sum >= 10 {
			carry = 1
		}
		newNode := &ListNode{Val: sum % 10}
		node.Next = newNode
		node = node.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + carry
		carry = 0
		if sum >= 10 {
			carry = 1
		}
		newNode := &ListNode{Val: sum % 10}
		node.Next = newNode
		node = node.Next
		l2 = l2.Next
	}
	if carry == 1 {
		newNode := &ListNode{Val: 1}
		node.Next = newNode
		node = node.Next
	}
	return reverseListNode(dummy.Next)
}

func reverseListNode(head *ListNode) *ListNode {
	var prev *ListNode
	var curr *ListNode
	curr = head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}
