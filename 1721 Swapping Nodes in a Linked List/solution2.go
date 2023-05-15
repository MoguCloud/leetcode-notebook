// Your runtime beats 73.97 % of golang submissions (168 ms)
// Your memory usage beats 23.29 % of golang submissions (10.8 MB)
//
// 双指针
// 节点交换
// 需要分别处理2个节点连续和不连续的情况
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func swapNodes(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	node := dummy
	for i := 1; i < k; i++ {
		node = node.Next
	}
	beginPrev := node
	endPrev := dummy
	node = node.Next
	for node.Next != nil {
		node = node.Next
		endPrev = endPrev.Next
	}
	if endPrev == beginPrev {
		return dummy.Next
	}
	if beginPrev.Next == endPrev {
		begin := beginPrev.Next
		end := endPrev.Next
		begin.Next = end.Next
		end.Next = begin
		beginPrev.Next = end
	} else if endPrev.Next == beginPrev {
		begin := beginPrev.Next
		end := endPrev.Next
		end.Next = begin.Next
		begin.Next = end
		endPrev.Next = begin
	} else {
		beginPrev.Next.Next, endPrev.Next.Next = endPrev.Next.Next, beginPrev.Next.Next
		beginPrev.Next, endPrev.Next = endPrev.Next, beginPrev.Next
	}

	return dummy.Next
}
