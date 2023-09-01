// Your runtime beats 82.55 % of golang submissions (12 ms)
// Your memory usage beats 32.21 % of golang submissions (7.52 MB)
//
// BFS
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

func nearestExit(maze [][]byte, entrance []int) int {
	visited := make([][]bool, len(maze))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(maze[0]))
	}
	q := [][2]int{[2]int{entrance[0], entrance[1]}}
	ret := 0
	for len(q) > 0 {
		newQ := [][2]int{}
		for _, p := range q {
			if p[0] < 0 || p[0] >= len(maze) || p[1] < 0 || p[1] >= len(maze[0]) {
				continue
			}
			if visited[p[0]][p[1]] {
				continue
			}
			if maze[p[0]][p[1]] == '+' {
				continue
			}
			if (p[0] == 0 || p[0] == len(maze)-1 || p[1] == 0 || p[1] == len(maze[0])-1) && !(p[0] == entrance[0] && p[1] == entrance[1]) {
				return ret
			}
			visited[p[0]][p[1]] = true
			newQ = append(newQ, [2]int{p[0] + 1, p[1]})
			newQ = append(newQ, [2]int{p[0] - 1, p[1]})
			newQ = append(newQ, [2]int{p[0], p[1] + 1})
			newQ = append(newQ, [2]int{p[0], p[1] - 1})
		}
		ret++
		q = newQ
	}
	return -1
}
