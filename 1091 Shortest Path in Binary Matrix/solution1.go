// Time Limit Exceeded
//
// DFS

type point struct {
	x int
	y int
}

var unreachable int = 10001

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid)-1] == 1 {
		return -1
	}
	memo := make(map[point]struct{})
	ret := dfs(grid, 0, 0, 1, &memo)
	if ret == unreachable {
		ret = -1
	}
	return ret
}

func dfs(grid [][]int, x int, y int, path int, memo *map[point]struct{}) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid) {
		return unreachable
	}
	if _, ok := (*memo)[point{x, y}]; ok {
		return unreachable
	}
	if x == len(grid)-1 && y == len(grid)-1 {
		return path
	}
	(*memo)[point{x, y}] = struct{}{}
	if grid[x][y] == 1 {
		return unreachable
	}
	paths := []int{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			paths = append(paths, dfs(grid, x+i, y+j, path+1, memo))
		}
	}
	delete(*memo, point{x, y})
	return min(paths)
}

func min(nums []int) int {
	m := nums[0]
	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}
