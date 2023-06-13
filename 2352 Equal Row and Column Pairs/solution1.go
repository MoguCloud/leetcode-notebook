// Your runtime beats 56.76 % of golang submissions (116 ms)
// Your memory usage beats 81.98 % of golang submissions (7.5 MB)
//
// Brute Force
//
// 时间复杂度 O(n^3)
// 空间复杂度 O(1)

func equalPairs(grid [][]int) int {
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			eq := true
			for k := 0; k < len(grid); k++ {
				if grid[i][k] != grid[k][j] {
					eq = false
					break
				}
			}
			if eq {
				ret++
			}
		}
	}
	return ret
}
