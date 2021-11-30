// Your runtime beats 28.68 % of golang submissions (96 ms)
// Your memory usage beats 48.06 % of golang submissions (7.7 MB)
//
// BFS
//
// bfs 最先找到的路径就是最短路径

type point struct {
	x int
	y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid)-1] == 1 {
		return -1
	}
	queue := []point{point{0, 0}}
	distance := 0
	visited := make(map[point]struct{})
	for len(queue) > 0 {
		newQueue := []point{}
		distance += 1
		for _, p := range queue {
			if p.x == len(grid)-1 && p.y == len(grid)-1 {
				return distance
			}
			if p.x < 0 || p.x >= len(grid) || p.y < 0 || p.y >= len(grid) {
				continue
			}
			if grid[p.x][p.y] == 1 {
				continue
			}
			if _, ok := visited[point{p.x, p.y}]; ok {
				continue
			} else {
				visited[point{p.x, p.y}] = struct{}{}
			}
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					newQueue = append(newQueue, point{p.x + i, p.y + j})
				}
			}
		}
		queue = newQueue
	}
	return -1
}
