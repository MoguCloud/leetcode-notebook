// Your runtime beats 64.52 % of golang submissions (4 ms)
// Your memory usage beats 49.46 % of golang submissions (4.3 MB)
//
// 中序遍历，先访问左节点，再访问根节点，最后访问右节点
// 后序遍历，先访问左节点，再访问右节点，最后访问根节点
// 中序遍历结果 [左子树，根节点，右子树]，其中左子树、右子树展开也是[左子树，根节点，右子树]，同样左右子树也是中序遍历
// 后序遍历结果 [左子树，右子树，根节点]，其中左子树、右子树展开也是[左子树，右子树，根节点]，同样左右子树也是后序遍历
// 可以看出后序遍历，最后一个节点就是根节点
// 根据根节点可以将中序遍历结果拆分成左子树的中序遍历和右子树的中序遍历
// 再根据左子树的中序遍历数组长度和右子树的中序遍历数组长度将后序遍历数组拆成左子树的后序遍历和右子树后序遍历
// 再根据左子树的中序遍历和后序遍历找出左子树和根节点，根据右子树的中序遍历和后序遍历找出右子树的根节点
// 直到遍历数组为空
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) <= 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootIdx := 0
	for rootIdx = 0; rootIdx < len(inorder); rootIdx++ {
		if inorder[rootIdx] == root.Val {
			break
		}
	}
	leftInorder := inorder[:rootIdx]
	rightInorder := inorder[rootIdx+1:]
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder) : len(leftInorder)+len(rightInorder)]
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}
