// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 23.68 % of golang submissions (2.2 MB)
//
// 前序遍历
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, tmp int) int {
	if root == nil {
		return 0
	}
	tmp = tmp*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return tmp
	}
	return dfs(root.Right, tmp) + dfs(root.Left, tmp)
}
