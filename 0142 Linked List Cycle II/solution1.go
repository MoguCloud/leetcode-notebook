// Your runtime beats 95.27 % of golang submissions (3 ms)
// Your memory usage beats 16.99 % of golang submissions (5 MB)
//
// Hash
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func detectCycle(head *ListNode) *ListNode {
	memo := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := memo[head]; ok {
			return head
		}
		memo[head] = struct{}{}
		head = head.Next
	}
	return nil
}
