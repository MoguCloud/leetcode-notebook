// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (3.7 MB)
//
// 层次遍历，通过修改节点的 Val 来存节点的序号
// 每次入队时修改子节点的 Val，左子节点的 Val 为父节点的 Val * 2 + 1，右子节点的 Val 为父节点的 Val * 2 + 2
// 层次的宽度就为 该层最右的节点的序号 - 该层最左的节点的序号 + 1

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	root.Val = 0
	queue = append(queue, root)
	maxWidth := 0
	for len(queue) != 0 {
		newQueue := []*TreeNode{}
		width := queue[len(queue)-1].Val - queue[0].Val + 1
		if width > maxWidth {
			maxWidth = width
		}
		for _, node := range queue {
			if node.Left != nil {
				node.Left.Val = 2*node.Val + 1
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 2
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}
	return maxWidth
}
