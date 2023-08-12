// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 8.18 % of golang submissions (3.2 MB)
//
// DFS
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

// 记忆化
var memo map[[2]int]int

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	memo = make(map[[2]int]int)
	ret := dfs(obstacleGrid, 0, 0)
	return ret
}

func dfs(obstacleGrid [][]int, i int, j int) int {
	memoKey := [2]int{i, j}
	ret, ok := memo[memoKey]
	if ok {
		return ret
	}
	if i >= len(obstacleGrid) || i < 0 || j >= len(obstacleGrid[0]) || j < 0 || obstacleGrid[i][j] == 1 {
		// 出边界或者有障碍物
		ret = 0
	} else if i == len(obstacleGrid)-1 && j == len(obstacleGrid[0])-1 && obstacleGrid[i][j] == 0 {
		// 到达终点
		ret = 1
	} else {
		// 向下DFS
		ret += dfs(obstacleGrid, i+1, j)
		// 向右DFS
		ret += dfs(obstacleGrid, i, j+1)
	}
	memo[memoKey] = ret
	return ret
}
