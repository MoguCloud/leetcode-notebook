// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (3.3 MB)
//
// 后序遍历
//
// 变量 count 需要移动的次数
// 变量 moreCoins 当前节点比 1 多出来的硬币数量
// 当遍历到叶子节点时，root.Val - 1 即为多出来的硬币数量(moreCoins) (可能为正，可能为负)
// 即 当前节点需要给父节点 moreCoins 枚硬币，也就是需要移动 |moreCoins| 次硬币，当前节点才能变成1 (因为可能为负)
// 当左子节点和右子节点都移动完变成 1 后，当前节点的硬币数量为 root.Val + 左子节点的 moreCoins + 右子节点的 moreCoins
// 当一个节点的左子节点和右子节点硬币都为 1 时，该节点就可以视为叶子节点，只用关心当前节点的硬币数量即可
// 所以是需要做后序遍历
//
// 时间复杂度 O(n)
// 空间复杂度 O(log n)

var count int

func distributeCoins(root *TreeNode) int {
	count = 0
	dfs(root)
	return count
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	root.Val += dfs(root.Left) + dfs(root.Right)
	moreCoins := root.Val - 1
	count += abs(moreCoins)
	return moreCoins
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
