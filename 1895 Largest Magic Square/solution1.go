// Your runtime beats 100 % of golang submissions (4 ms)
// Your memory usage beats 25.71 % of golang submissions (5.8 MB)
//
// 前缀和
// 分别求每个元素在行、列、对角线上的前缀和
// 求完前缀和之后，幻方尺寸从min(m, n)到0开始遍历，当满足幻方条件时就是最大的幻方尺寸
// 前缀和数组size为[m+2, n+2]，可以避免判断第1个元素和最后个元素的边界条件判断
//
// 时间复杂度 O(min(m,n)^2*m*n)
// 空间复杂度 O(m*n)

func largestMagicSquare(grid [][]int) int {
	rowSum := make([][]int, len(grid)+2) // 行前缀和
	colSum := make([][]int, len(grid)+2) // 列前缀和
	diSum1 := make([][]int, len(grid)+2) // 左上到右下对角线
	diSum2 := make([][]int, len(grid)+2) // 右上到走下对角线
	// 前缀和数组初始化，避免边界条件判断
	for i := 0; i < len(rowSum); i++ {
		rowSum[i] = make([]int, len(grid[0])+2)
		colSum[i] = make([]int, len(grid[0])+2)
		diSum1[i] = make([]int, len(grid[0])+2)
		diSum2[i] = make([]int, len(grid[0])+2)
	}
	// 前缀和数组计算
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			rowSum[i+1][j+1] = grid[i][j] + rowSum[i+1][j+2]
			colSum[i+1][j+1] = grid[i][j] + colSum[i+2][j+1]
			diSum1[i+1][j+1] = grid[i][j] + diSum1[i+2][j+2]
		}
		for j := 0; j < len(grid[0]); j++ {
			diSum2[i+1][j+1] = grid[i][j] + diSum2[i+2][j]
		}
	}
	m := len(grid)
	n := len(grid[0])
	s := min(m, n)
OutLoop:
	// size 从大到小遍历，第一次满足条件即可返回
	for ; s > 1; s-- {
		for i := 0; i < m; i++ {
			if i+s > m {
				continue
			}
		Loop:
			for j := 0; j < n; j++ {
				if j+s > n {
					continue Loop
				}
				var col, row, di1, di2 int
				di1 = diSum1[i+1][j+1] - diSum1[i+1+s][j+1+s]
				di2 = diSum2[i+1][j+s] - diSum2[i+1+s][j]
				if di1 != di2 {
					continue Loop
				}
				for ss := 0; ss < s; ss++ {
					col = colSum[i+1][j+1+ss] - colSum[i+1+s][j+1+ss]
					if col != di1 {
						continue Loop
					}
					row = rowSum[i+1+ss][j+1] - rowSum[i+1+ss][j+1+s]
					if row != di1 {
						continue Loop
					}
				}
				break OutLoop
			}
		}
	}
	return s
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
