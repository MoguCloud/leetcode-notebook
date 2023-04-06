// Your runtime beats 80.85 % of golang submissions (8 ms)
// Your memory usage beats 55.32 % of golang submissions (4.7 MB)
//
// DFS
// 因为最外面一圈的土地无法构成封闭岛屿，所以将最外圈的土地变成水，然后计算封闭岛屿即可
//
// 时间复杂度 O(n*m)
// 空间复杂度 O(n*m)

func closedIsland(grid [][]int) int {
	ret := 0
	// 将最外圈土地变成水
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 0 {
			dfs(grid, i, 0)
		}
		if grid[i][len(grid[0])-1] == 0 {
			dfs(grid, i, len(grid[0])-1)
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 0 {
			dfs(grid, 0, i)
		}
		if grid[len(grid)-1][i] == 0 {
			dfs(grid, len(grid)-1, i)
		}
	}
	// 计算封闭岛屿
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 0 {
				ret++
				dfs(grid, i, j)
			}
		}
	}
	return ret
}

func dfs(grid [][]int, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
