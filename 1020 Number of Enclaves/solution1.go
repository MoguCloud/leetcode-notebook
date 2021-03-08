// DFS
//
// Your runtime beats 83.33 % of golang submissions
// Your memory usage beats 75 % of golang submissions (6.5 MB)
//
// 图最外面一圈进行DFS，将相邻的 1 改成 0。DFS完成后，统计剩余的 1 的数量。

func numEnclaves(grid [][]int) int {
	for i := 0; i < len(grid[0]); i++ {
		dfs(grid, 0, i)
		dfs(grid, len(grid)-1, i)
	}
	for i := 0; i < len(grid); i++ {
		dfs(grid, i, 0)
		dfs(grid, i, len(grid[0])-1)
	}
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ret += 1
			}
		}
	}
	return ret
}

func dfs(grid [][]int, x int, y int) {
	if grid[x][y] == 1 {
		grid[x][y] = 0
		if x+1 >= 0 && x+1 <= len(grid)-1 {
			dfs(grid, x+1, y)
		}
		if x-1 >= 0 && x-1 <= len(grid)-1 {
			dfs(grid, x-1, y)
		}
		if y+1 >= 0 && y+1 <= len(grid[0])-1 {
			dfs(grid, x, y+1)
		}
		if y-1 >= 0 && y-1 <= len(grid[0])-1 {
			dfs(grid, x, y-1)
		}
	}
}
