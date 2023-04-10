// Your runtime beats 66.44 % of golang submissions (2 ms)
// Your memory usage beats 80.27 % of golang submissions (2.3 MB)
//
// 原地旋转
//
// 时间复杂度 O(n*n)
// 空间复杂度 O(1)

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := i; j < n-1-i; j++ {
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
	return
}
