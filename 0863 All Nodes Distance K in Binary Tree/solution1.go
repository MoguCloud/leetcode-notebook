// Your runtime beats 24.56 % of golang submissions (5 ms)
// Your memory usage beats 35.9 % of golang submissions (3.5 MB)
//
// DFS
// 从 target 节点向子节点以及父节点进行遍历距离k，遍历时记录已经访问过的节点，避免重复遍历
// 先进行一次 DFS，使用 hash 表记录每个节点的父节点
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}
	// 存储每个节点的父节点的 hash 表
	parentMap := make(map[int]*TreeNode)
	var findParent func(*TreeNode)
	// 记录父节点
	findParent = func(node *TreeNode) {
		if node.Left != nil {
			parentMap[node.Left.Val] = node
			findParent(node.Left)
		}
		if node.Right != nil {
			parentMap[node.Right.Val] = node
			findParent(node.Right)
		}
	}
	findParent(root)
	// 记录已经遍历过的节点，防止重复便利
	visited := make(map[int]struct{})
	ret := []int{}
	var findAns func(*TreeNode, int)
	// 从 target 节点向子节点以及父节点遍历距离 k
	findAns = func(node *TreeNode, deep int) {
		// 避免重复遍历
		if _, ok := visited[node.Val]; ok {
			return
		}
		visited[node.Val] = struct{}{}
		// 到达距离 k，进行记录
		if deep == 0 && node != target {
			ret = append(ret, node.Val)
		}
		// 左子节点
		if node.Left != nil {
			findAns(node.Left, deep-1)
		}
		// 右子节点
		if node.Right != nil {
			findAns(node.Right, deep-1)
		}
		// 父节点
		if parent, ok := parentMap[node.Val]; ok {
			findAns(parent, deep-1)
		}
	}
	findAns(target, k)
	return ret
}
