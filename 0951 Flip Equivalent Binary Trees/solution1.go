// Recursion
//
// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 27.59 % of golang submissions (2.5 MB)
//
// 递归判断当前节点是否相同，相同再判断左右孩子是否相同。
// 需要注意 root1 和 root2 的左右孩子的值可能都相同，存在需要翻转再判断的情况。

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	if root1 == nil && root2 == nil {
		return true
	}
	if root1.Val != root2.Val {
		return false
	} else {
		return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) || (flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
	}

}
