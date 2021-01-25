// 层次遍历
//
// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 97.45 % of golang submissions (5.9 MB)
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		newQueue := []*Node{}
		for i := 0; i < len(queue); i++ {
			if i != len(queue) - 1 {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}
		queue = newQueue
	}
	return root
}
