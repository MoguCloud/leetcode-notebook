// Your runtime beats 93.4 % of golang submissions (8 ms)
// Your memory usage beats 62.40 % of golang submissions (5.1 MB)
//
// DFS
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

var max int

func maxAreaOfIsland(grid [][]int) int {
	max = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j, 0)
			}
		}
	}
	return max
}

func dfs(grid [][]int, x int, y int, area int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != 1 {
		return area
	}
	area = area + 1
	if area > max {
		max = area
	}
	grid[x][y] = 0
	area = dfs(grid, x+1, y, area)
	area = dfs(grid, x-1, y, area)
	area = dfs(grid, x, y+1, area)
	area = dfs(grid, x, y-1, area)
	return area
}
