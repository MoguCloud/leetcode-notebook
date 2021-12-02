// Your runtime beats 87.3 % of golang submissions (8 ms)
// Your memory usage beats 39.68 % of golang submissions (6.3 MB)
//
// 二维前缀和
// https://pic.leetcode-cn.com/1611719222-caavDU-image.png
// prefixSum[x][y] 为 mat矩阵从(0, 0) 到 (x-1, y-1)形成的矩阵的和
// (x1, y1), (x2, y2) 形成的矩阵和为 prefixSum[x2+1][y2+1] - prefixSum[x2+1][y1] - prefixSum[x1][y2+1] + prefixSum[x1][y1]
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

func matrixBlockSum(mat [][]int, k int) [][]int {
	answer := make([][]int, len(mat))
	for i := 0; i < len(answer); i++ {
		answer[i] = make([]int, len(mat[i]))
	}
	prefixSum := make([][]int, len(mat)+1)
	for i := 0; i < len(prefixSum); i++ {
		prefixSum[i] = make([]int, len(mat[0])+1)
	}
	for i := 1; i < len(prefixSum); i++ {
		for j := 1; j < len(prefixSum[i]); j++ {
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + mat[i-1][j-1]
		}
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			x1 := max(0, i-k)
			x2 := min(i+k, len(mat)-1)
			y1 := max(0, j-k)
			y2 := min(j+k, len(mat[i])-1)
			answer[i][j] = prefixSum[x2+1][y2+1] - prefixSum[x2+1][y1] - prefixSum[x1][y2+1] + prefixSum[x1][y1]
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
