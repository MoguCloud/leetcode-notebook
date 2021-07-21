// Your runtime beats 96 % of golang submissions (20 ms)
// Your memory usage beats 49.14 % of golang submissions (6.9 MB)
//
// 将matrix分成 正方形的数组 和 长方形的数组两部分进行处理
// 如 [[1, 2, 3],
//     [4, 5, 6]]
// 可以拆分成 [[1, 2],  和 [[3],  两部分
//            [4, 5]]      [6]]
// 假设 matrix[m][n], m < n
// 正方形部分square[m][m]
// 从(0, 0)到(m, m)遍历正方形的对角线
// 如果square[i][i] < target，则target不在第i行
// 如果square[i][i] >= target，则target可能在第i行或者第i列
// 在第i行或者第i列进行查找时，从右往左/从下往上进行查找或者使用二分查找
// 如果在正方形部分没有找到，则在长方形部分进行查找
// 按长方形的最右边往下遍历或者最下边往右遍历num[i]
// 如果num[i] >= target，则target可能在第i行/列

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < min(m, n); i++ {
		if matrix[i][i] >= target {
			for j := i; j >= 0; j-- {
				if matrix[i][j] == target {
					return true
				} else if matrix[i][j] < target {
					break
				}
			}
			for j := i; j >= 0; j-- {
				if matrix[j][i] == target {
					return true
				} else if matrix[j][i] < target {
					break
				}
			}
		}
	}
	if m > n {
		for i := n; i < m; i++ {
			if matrix[i][n-1] >= target {
				for j := n - 1; j >= 0; j-- {
					if matrix[i][j] == target {
						return true
					}
					if matrix[i][j] < target {
						break
					}
				}
			}
		}
	}
	if m < n {
		for i := m; i < n; i++ {
			if matrix[m-1][i] >= target {
				for j := m - 1; j >= 0; j-- {
					if matrix[j][i] == target {
						return true
					}
					if matrix[j][i] < target {
						break
					}
				}
			}
		}
	}
	return false
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
