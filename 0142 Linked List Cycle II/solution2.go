// Your runtime beats 77.91 % of golang submissions (4 ms)
// Your memory usage beats 43.4 % of golang submissions (3.8 MB)
//
// 双指针
// 从head走到环入口长为a，环长度为b
// f = 2s
// f = s + nb
// => s = nb
// s再走a步可以到环入口
// 从head走a步可以到环入口
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for {
		if fast == slow {
			return slow
		}
		fast = fast.Next
		slow = slow.Next
	}
}
