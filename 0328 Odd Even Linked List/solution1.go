/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 32.86 % of golang submissions (3.3 MB)
//
// 将原来的列表分成奇数位链表和偶数位链表
// 然后将奇数位链表的最后一个元素的next指向偶数位链表的表头即可
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	oddDummy := &ListNode{}
	evenDummy := &ListNode{}
	oddNode := oddDummy
	evenNode := evenDummy
	node := head
	prev := &ListNode{}
	for i := 1; node != nil; i++ {
		prev.Next = nil
		if i&1 == 1 {
			oddNode.Next = node
			oddNode = oddNode.Next
		} else {
			evenNode.Next = node
			evenNode = evenNode.Next
		}
		prev = node
		node = node.Next
	}
	oddNode.Next = evenDummy.Next
	return oddDummy.Next
}
