// 递归
//
// Your runtime beats 81.48 % of golang submissions (72 ms)
// Your memory usage beats 85.19 % of golang submissions (6.9 MB)

var maxDepth int

func longestUnivaluePath(root *TreeNode) int {
	maxDepth = 0
	helper(root, 0)
	return maxDepth
}

func helper(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	if root.Left != nil && root.Left.Val == root.Val && root.Right != nil && root.Right.Val == root.Val {
		left := helper(root.Left, 1)
		right := helper(root.Right, 1)
		maxDepth = max(maxDepth, left+depth, right+depth, left+right)
		return max(left, right) + depth
	}
	if root.Left != nil {
		if root.Left.Val == root.Val {
			depth = helper(root.Left, depth+1)
			maxDepth = max(depth, maxDepth)
		} else {
			helper(root.Left, 0)
		}
	}
	if root.Right != nil {
		if root.Right.Val == root.Val {
			depth = helper(root.Right, depth+1)
			maxDepth = max(depth, maxDepth)
		} else {
			helper(root.Right, 0)
		}
	}
	return depth
}

func max(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}
