// Your runtime beats 28.57 % of golang submissions (33 ms)
// Your memory usage beats 91.48 % of golang submissions (7.4 MB)
//
// 层次遍历
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func sumEvenGrandparent(root *TreeNode) int {
	ret := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		newQ := []*TreeNode{}
		for _, n := range q {
			if n.Val&1 == 0 {
				ret += getVal(n)
			}
			if n.Left != nil {
				newQ = append(newQ, n.Left)
			}
			if n.Right != nil {
				newQ = append(newQ, n.Right)
			}
		}
		q = newQ
	}
	return ret
}

func getVal(root *TreeNode) int {
	ret := 0
	if root.Left != nil {
		if root.Left.Left != nil {
			ret += root.Left.Left.Val
		}
		if root.Left.Right != nil {
			ret += root.Left.Right.Val
		}
	}
	if root.Right != nil {
		if root.Right.Left != nil {
			ret += root.Right.Left.Val
		}
		if root.Right.Right != nil {
			ret += root.Right.Right.Val
		}
	}
	return ret
}
