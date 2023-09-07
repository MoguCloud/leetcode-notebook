// Your runtime beats 62.50 % of golang submissions (7 ms)
// Your memory usage beats 62.50 % of golang submissions (5.34 MB)
//
// DFS
// 后序遍历，先遍历叶子节点，后遍历根节点
// 遍历时候维护当前节点深度和最大深度
// 当到叶子节点时，深度等于最大深度，将返回值置为叶子节点
// 遍历完叶子节点后遍历根节点，如果根节点和左子树深度=右子树深度，且等于最大深度时
// 说明当前节点是最近公共祖先，将返回值置为当前节点
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	ret := root
	maxDepth := 0
	var dfs func(root *TreeNode, depth int) int
	dfs = func(root *TreeNode, depth int) int {
		if root.Left == nil && root.Right == nil {
			maxDepth = max(maxDepth, depth)
			if depth >= maxDepth {
				ret = root
			}
			return depth
		}
		leftDepth, rightDepth := 0, 0
		if root.Left != nil {
			leftDepth = dfs(root.Left, depth+1)
		}
		if root.Right != nil {
			rightDepth = dfs(root.Right, depth+1)
		}
		if leftDepth == rightDepth && leftDepth >= maxDepth {
			maxDepth = leftDepth
			ret = root
		}
		return max(rightDepth, leftDepth)
	}
	dfs(root, 0)
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
