// Your runtime beats 70 % of golang submissions (117 ms)
// Your memory usage beats 90 % of golang submissions (12.6 MB)
//
// DFS
// 如果从父节点过来的方向和去子节点的方向相同的话，长度从0开始算
// 如果从父节点过来的方向和去子节点的方向不同的话，长度加1
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

var ret int

func longestZigZag(root *TreeNode) int {
	ret = 0
	dfs(root, 0, 0)
	dfs(root, 1, 0)
	return ret
}

func dfs(root *TreeNode, direct int, distance int) {
	if root == nil {
		return
	}
	if distance > ret {
		ret = distance
	}
	if direct == 0 {
		dfs(root.Left, 0, 0)
		dfs(root.Left, 1, distance+1)
	}
	if direct == 1 {
		dfs(root.Right, 1, 0)
		dfs(root.Right, 0, distance+1)
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
