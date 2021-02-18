// BFS
//
// Your runtime beats 5.88 % of golang submissions (100 ms)
// Your memory usage beats 41.18 % of golang submissions (6.6 MB)

type pair struct {
	x int
	y int
}

func countServers(grid [][]int) int {
	gridSet := make(map[pair]struct{})
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				p := pair{x: i, y: j}
				gridSet[p] = struct{}{}
			}
		}
	}
	count := 0
	for k := range gridSet {
		c := bfs(gridSet, k, 1)
		if c > 1 {
			count += c
		}
	}
	return count
}

func bfs(gridSet map[pair]struct{}, p pair, count int) int {
	delete(gridSet, p)
	gridArr := []pair{}
	for k := range gridSet {
		if k.x == p.x || k.y == p.y {
			gridArr = append(gridArr, k)
			count += 1
		}
	}
	for i := 0; i < len(gridArr); i++ {
		delete(gridSet, gridArr[i])
	}
	for i := 0; i < len(gridArr); i++ {
		count = bfs(gridSet, gridArr[i], count)
	}
	return count
}
