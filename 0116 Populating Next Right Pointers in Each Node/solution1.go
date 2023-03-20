// Your runtime beats 43.82 % of golang submissions (8 ms)
// Your memory usage beats 13.78 % of golang submissions (7 MB)
//
// 层次遍历
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) != 0 {
		newQueue := []*Node{}
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
			if i != len(queue)-1 {
				node.Next = queue[i+1]
			}
		}
		queue = newQueue
	}
	return root
}
