// Your runtime beats 80.57 % of golang submissions (4 ms)
// Your memory usage beats 75.27 % of golang submissions (6.5 MB)
//
// 使用next指针代替队列进行层次遍历
// 记录每层的左边节点，使用该节点的next指针就可以遍历该层
// 每个节点将它的Left节点的Next指向Right节点，Right节点的Next指向该节点Next节点的Left节点
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func connect(root *Node) *Node {
	prevLeft := root
	for prevLeft != nil && prevLeft.Left != nil {
		for node := prevLeft; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
		prevLeft = prevLeft.Left
	}
	return root
}
