// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 90.75 % of golang submissions (2 MB)
//
// 模拟螺旋
//
// 时间复杂度 O(n*m)
// 空间复杂度 O(1)

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	ret := make([]int, m*n)
	// 剩余需要遍历的边界
	minM := 0
	maxM := m - 1
	minN := 0
	maxN := n - 1
	// 方向 0向右 1向下 2向左 3向上
	dir := 0
	// 当前坐标
	i := 0
	j := 0
	for idx := 0; idx < m*n; idx++ {
		ret[idx] = matrix[i][j]
		switch dir % 4 {
		case 0:
			// 向右遍历，当j<右边界时，向右；到达边界时，调整方向开始向下遍历，同时上边界向下移动
			if j < maxN {
				j++
			} else {
				dir++
				i++
				minM++
			}
		case 1:
			// 向下遍历，当i<下边界时，向下；到达边界时，调整方向开始向左遍历，同时右边界向左移动
			if i < maxM {
				i++
			} else {
				dir++
				j--
				maxN--
			}
		case 2:
			// 向左遍历，当j>左边界时，向左；到达边界时，调整方向开始向上遍历，同时下边界向上移动
			if j > minN {
				j--
			} else {
				dir++
				i--
				maxM--
			}
		case 3:
			// 向上遍历，当j>上边界时，向上；到达边界时，调整方向开始向右遍历，同时左边界向右移动
			if i > minM {
				i--
			} else {
				dir++
				j++
				minN++
			}
		}
	}
	return ret
}
