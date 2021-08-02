// Your runtime beats 52.63 % of golang submissions (48 ms)
// Your memory usage beats 57.89 % of golang submissions (7.2 MB)
//
// 二叉树满足
// 只有1个根节点 入度=0
// 其他节点 入度=1 出度<=2
// 不成环 => 从根节点遍历一遍，所有节点都只被遍历到一次
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	inDegree := make([]int, n)
	outDegree := make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			outDegree[i] += 1
			inDegree[leftChild[i]] += 1
		}
		if rightChild[i] != -1 {
			outDegree[i] += 1
			inDegree[rightChild[i]] += 1
		}
	}
	root := -1
	for i := 0; i < n; i++ {
		if inDegree[i] > 1 {
			return false
		} else if inDegree[i] == 0 {
			if root != -1 {
				return false
			}
			root = i
		}
		if outDegree[i] > 2 {
			return false
		}
	}
	if root == -1 {
		return false
	}
	memo := make(map[int]bool)
	stack := []int{}
	stack = append(stack, root)
	for len(stack) > 0 {
		root := stack[0]
		stack = stack[1:]
		left := leftChild[root]
		if left != -1 {
			if _, ok := memo[left]; ok {
				return false
			}
			memo[left] = true
			stack = append(stack, left)
		}
		right := rightChild[root]
		if right != -1 {
			if _, ok := memo[right]; ok {
				return false
			}
			memo[right] = true
			stack = append(stack, right)
		}
	}
	memo[root] = true
	for i := 0; i < n; i++ {
		if _, ok := memo[i]; !ok {
			return false
		}
	}
	return true
}
