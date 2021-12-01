// Your runtime beats 89.58 % of golang submissions (16 ms)
// Your memory usage beats 77.08 % of golang submissions (9.3 MB)
//
// Recursion
//
// 递归实现构造满二叉树
// n个节点的二叉树可以拆分为
//   1个节点用于根节点
//   i个节点用于构造左子树
//   n-i-1个节点用于构造右子树
// (i, n-i-1都为奇数)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var memo = make(map[int][]*TreeNode)

func allPossibleFBT(n int) []*TreeNode {
	ret := []*TreeNode{}
	// 节点数量为偶数时，不能构造出满二叉树
	if n&1 == 0 {
		return ret
	}
	// 通过记忆进行优化
	if arr, ok := memo[n]; ok {
		return arr
	}
	if n == 1 {
		ret = append(ret, &TreeNode{Val: 0})
	}
	// n为1时不会进入该循环，n为1时需要特殊处理
	for i := 1; n-i-1 >= 0; i += 2 {
		// 构造出所有的左子树
		leftTrees := allPossibleFBT(i)
		// 构造出所有的右子树
		rightTrees := allPossibleFBT(n - i - 1)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				// 构造1个根节点，并与左子树、右子树进行组合
				root := &TreeNode{
					Val:   0,
					Left:  leftTree,
					Right: rightTree,
				}
				ret = append(ret, root)
			}
		}
	}
	memo[n] = ret
	return ret
}
