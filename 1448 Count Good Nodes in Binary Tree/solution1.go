// DFS
//
// Your runtime beats 95.35 % of golang submissions (104 ms)
// Your memory usage beats 83.72 % of golang submissions (10.9 MB)

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val, 0)
}

func dfs(root *TreeNode, max int, count int) int {
	if root.Val >= max {
		count += 1
		max = root.Val
	}
	if root.Left != nil {
		count = dfs(root.Left, max, count)
	}
	if root.Right != nil {
		count = dfs(root.Right, max, count)
	}
	return count
}
