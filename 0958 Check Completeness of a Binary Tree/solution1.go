// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 40.91 % of golang submissions (3.2 MB)
//
// 层次遍历
// 当出现空节点时候则后续不允许出现非空节点
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		newQueue := []*TreeNode{}
		for i, node := range queue {
			if node == nil {
				for j := i + 1; j < len(queue); j++ {
					if queue[j] != nil {
						return false
					}
				}
				for j := 0; j < len(newQueue); j++ {
					if newQueue[j] != nil {
						return false
					}
				}
				continue
			}
			newQueue = append(newQueue, node.Left)
			newQueue = append(newQueue, node.Right)
		}
		queue = newQueue
	}
	return true
}
