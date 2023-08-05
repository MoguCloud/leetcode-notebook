// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 90.83 % of golang submissions (4.18 MB)
//
// 记忆化递归
// [0..n] 对于 i (0 <= i <= n) 作为根节点要构成二叉搜索树，可以通过递归生成，i 作为根节点， [0..i-1]递归生成i的左子节点，[i+1..n]递归生成i的右子节点

package main

func generateTrees(n int) []*TreeNode {
	// 用于记忆化搜索，避免重复计算
	memo := make(map[string][]*TreeNode)
	var genTrees func(start int, end int) []*TreeNode
	// genTrees 构造[start..end]的二叉搜索树
	genTrees = func(start int, end int) []*TreeNode {
		// 记忆化使用的key
		memoKey := fmt.Sprintf("%d-%d", start, end)
		// 已经计算过就直接返回
		if ans, ok := memo[memoKey]; ok {
			return ans
		}
		ret := []*TreeNode{}
		// 当start > end时，则不存在，因为后续需要同时遍历左子节点和右子节点来构造树，所以空树需要返回nil，否则range []*TreeNode{}时不会进入循环语句
		if start > end {
			ret = append(ret, nil)
			return ret
		}
		if start == end {
			root := &TreeNode{Val: start}
			ret = append(ret, root)
			return ret
		}
		for i := start; i <= end; i++ {
			// 构造左子树
			leftNodes := genTrees(start, i-1)
			// 构造右子树
			rightNodes := genTrees(i+1, end)
			// 遍历所有左子树和右子树，构造所有可能的情况
			for _, left := range leftNodes {
				for _, right := range rightNodes {
					root := &TreeNode{Val: i}
					root.Left = left
					root.Right = right
					ret = append(ret, root)
				}
			}
		}
		// 记录结果，避免后续重复计算
		memo[memoKey] = ret
		return ret
	}
	return genTrees(1, n)
}
