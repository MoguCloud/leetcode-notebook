// Your runtime beats 48.74 % of golang submissions (24 ms)
// Your memory usage beats 80.67 % of golang submissions (6.7 MB)
//
// 模拟
//
// 时间复杂度 O(n*m)
// 空间复杂度 O(1)

func findBall(grid [][]int) []int {
	ret := make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		ret[i] = dfs(grid, 0, i)
	}
	return ret
}

func dfs(grid [][]int, row int, col int) int {
	var ret int
	if col == -1 {
		ret = -1
	} else if row == len(grid) {
		// 掉落到终点
		ret = col
	} else if grid[row][col] == 1 && col+1 < len(grid[0]) && grid[row][col+1] == 1 {
		// 向右侧掉落
		ret = dfs(grid, row+1, col+1)
	} else if grid[row][col] == -1 && col-1 >= 0 && grid[row][col-1] == -1 {
		// 向左侧掉落
		ret = dfs(grid, row+1, col-1)
	} else {
		ret = -1
	}
	return ret
}
