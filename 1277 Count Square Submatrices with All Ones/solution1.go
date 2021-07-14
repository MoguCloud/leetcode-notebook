// Your runtime beats 41.38 % of golang submissions (84 ms)
// Your memory usage beats 8.62 % of golang submissions (8.1 MB)
//
// Brute-Force
//
// 从左到右，从上到下依次遍历，遍历到每个元素的时候，
// 依次判断是否能构成 1x1, 2x2 ... 的正方形
//
// 时间复杂度 O(M ^ 2 * N ^ 2)
// 空间复杂度 O(1)

func countSquares(matrix [][]int) int {
	ret := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				goto Loop
			}
			for k := 0; i+k < len(matrix) && j+k < len(matrix[i]); k++ {
				for kk := 0; kk <= k; kk++ {
					if matrix[i+k][j+kk] == 0 || matrix[i+kk][j+k] == 0 {
						goto Loop
					}
				}
				ret += 1
			}
		Loop:
		}
	}
	return ret
}
