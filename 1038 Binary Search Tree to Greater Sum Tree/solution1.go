// Your runtime beats 66.10 % of golang submissions (2 ms)
// Your memory usage beats 74.58 % of golang submissions (2.3 MB)
//
// 反中序遍历
// 右 -> 根 -> 左 进行遍历
// 当前节点的值 = 该节点右侧节点的和 + 当前节点的值
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func bstToGst(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func dfs(root *TreeNode, sum int) int {
	if root.Right != nil {
		sum = dfs(root.Right, sum)
	}
	sum += root.Val
	root.Val = sum
	if root.Left != nil {
		sum = dfs(root.Left, sum)
	}
	return sum
}
