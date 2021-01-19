// BFS
//
// Your runtime beats 78.48 % of golang submissions (28 ms)
// Your memory usage beats 81.01 % of golang submissions (7.1 MB)

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	sum := 0
	for len(q) > 0 {
		newQueue := []*TreeNode{}
		sum = 0
		for i := 0; i < len(q); i++ {
			sum += q[i].Val
			if q[i].Left != nil {
				newQueue = append(newQueue, q[i].Left)
			}
			if q[i].Right != nil {
				newQueue = append(newQueue, q[i].Right)
			}
		}
		q = newQueue
	}
	return sum
}
