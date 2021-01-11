// 递归 后序遍历
//	
// Your runtime beats 100 % of golang submissions (8 ms)
// Your memory usage beats 71.79 % of golang submissions (8.3 MB)

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
