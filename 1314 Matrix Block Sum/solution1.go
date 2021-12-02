// Your runtime beats 87.3 % of golang submissions (8 ms)
// Your memory usage beats 85.71 % of golang submissions (6.2 MB)
//
// 先计算mat[0][0]的矩阵和
// j != 0 的矩阵和 等于 当前行的前一项的矩阵和 减去前一项矩阵和第一列 再加上当前项的矩阵和最后一列
// j == 0 的矩阵和 等于 当前列的上一项的矩阵和 减去上一项矩阵的第一行 再加上当前项的矩阵和最后一行
//
// 时间复杂度 O(m*n*k)
// 空间复杂度 O(1)

func matrixBlockSum(mat [][]int, k int) [][]int {
	answer := make([][]int, len(mat))
	for i := 0; i < len(answer); i++ {
		answer[i] = make([]int, len(mat[i]))
	}
	sum := 0
	for r := 0; r <= min(k, len(mat)-1); r++ {
		for c := 0; c <= min(k, len(mat[r])-1); c++ {
			sum += mat[r][c]
		}
	}
	answer[0][0] = sum
	for i := 0; i < len(answer); i++ {
		for j := 0; j < len(answer[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if j == 0 {
				sum = answer[i-1][j]
				if i-k-1 >= 0 {
					for c := max(0, j-k); c <= min(len(mat[i])-1, j+k); c++ {
						sum -= mat[i-k-1][c]
					}
				}
				if i+k < len(mat) {
					for c := max(0, j-k); c <= min(len(mat[i])-1, j+k); c++ {
						sum += mat[i+k][c]
					}
				}

			} else {
				sum = answer[i][j-1]
				if j-k-1 >= 0 {
					for r := max(0, i-k); r <= min(len(mat)-1, i+k); r++ {
						sum -= mat[r][j-k-1]
					}
				}
				if j+k < len(mat[i]) {
					for r := max(0, i-k); r <= min(len(mat)-1, i+k); r++ {
						sum += mat[r][j+k]
					}
				}
			}
			answer[i][j] = sum
		}
	}
	return answer
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
