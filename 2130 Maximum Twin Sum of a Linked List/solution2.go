// Your runtime beats 59.9 % of golang submissions (196 ms)
// Your memory usage beats 89.61 % of golang submissions (9.2 MB)
//
// 使用快慢指针找到链表的中点，翻转后半段链表，再求和
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func pairSum(head *ListNode) int {
	// 找中点
	fast := head
	slow := head
	count := 0
	for fast != nil {
		fast = fast.Next.Next
		slow = slow.Next
		count++
	}
	// 翻转链表
	curr := slow
	next := curr.Next
	for next != nil {
		tmp := next.Next
		next.Next = curr
		curr = next
		next = tmp
	}
	ret := 0
	// 求和
	for ; count > 0; count-- {
		sum := curr.Val + head.Val
		if sum > ret {
			ret = sum
		}
		curr = curr.Next
		head = head.Next
	}
	return ret
}
