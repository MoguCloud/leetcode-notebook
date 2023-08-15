// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 17.26 % of golang submissions (2.42 MB)
//
// 模拟
// 将原链表分成一个全是小于特定值的链表A和一个全是大于等于特定值的链表B
// 再将A的链表尾指向B的链表头
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func partition(head *ListNode, x int) *ListNode {
	ltDummy := &ListNode{}
	geDummy := &ListNode{}
	lt := ltDummy
	ge := geDummy
	node := head
	for node != nil {
		if node.Val < x {
			lt.Next = node
			lt = lt.Next
		} else {
			ge.Next = node
			ge = ge.Next
		}
		node = node.Next
	}
	ge.Next = nil
	lt.Next = geDummy.Next
	return ltDummy.Next
}
