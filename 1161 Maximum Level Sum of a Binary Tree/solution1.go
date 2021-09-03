// Your runtime beats 89.58 % of golang submissions (112 ms)
// Your memory usage beats 91.67 % of golang submissions (8.3 MB)
//
// BFS 层次遍历
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func maxLevelSum(root *TreeNode) int {
	maxSum := root.Val
	maxLevel := 1
	nodeArr := []*TreeNode{}
	nodeArr = append(nodeArr, root)
	curLevel := 1
	for len(nodeArr) > 0 {
		newArr := []*TreeNode{}
		sum := 0
		for i := 0; i < len(nodeArr); i++ {
			sum += nodeArr[i].Val
			if nodeArr[i].Left != nil {
				newArr = append(newArr, nodeArr[i].Left)
			}
			if nodeArr[i].Right != nil {
				newArr = append(newArr, nodeArr[i].Right)
			}
		}
		if sum > maxSum {
			maxSum = sum
			maxLevel = curLevel
		}
		curLevel += 1
		nodeArr = newArr
	}
	return maxLevel
}
