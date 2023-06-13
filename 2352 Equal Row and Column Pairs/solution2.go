// Your runtime beats 33.33 % of golang submissions (168 ms)
// Your memory usage beats 74.77 % of golang submissions (7.6 MB)
//
// Hash
// 先遍历行，将行作为key存入hash table中
// 再遍历列，判断列是否存在于hash table中
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n^2)

func equalPairs(grid [][]int) int {
	ret := 0
	memo := make(map[string]int)
	for i := 0; i < len(grid); i++ {
		memo[row2key(grid, i)]++
	}
	for i := 0; i < len(grid); i++ {
		ret += memo[col2key(grid, i)]
	}
	return ret
}

func row2key(grid [][]int, row int) string {
	return fmt.Sprint(grid[row])
}

func col2key(grid [][]int, col int) string {
	column := make([]int, len(grid))
	for i := 0; i < len(grid); i++ {
		column[i] = grid[i][col]
	}
	return fmt.Sprint(column)
}
