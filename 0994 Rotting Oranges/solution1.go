// Your runtime beats 33.11 % of golang submissions (6 ms)
// Your memory usage beats 10.49 % of golang submissions (3.5 MB)
//
// BFS
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

func orangesRotting(grid [][]int) int {
	q := [][2]int{}
	fresh := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				q = append(q, [2]int{i + 1, j})
				q = append(q, [2]int{i - 1, j})
				q = append(q, [2]int{i, j + 1})
				q = append(q, [2]int{i, j - 1})
			}
		}
	}
	ret := 0
	for {
		if fresh == 0 {
			return ret
		}
		if len(q) == 0 {
			return -1
		}
		ret++
		newQ := [][2]int{}
		for k := 0; k < len(q); k++ {
			i := q[k][0]
			j := q[k][1]
			if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
				continue
			}
			fresh--
			grid[i][j] = 2
			newQ = append(newQ, [2]int{i + 1, j})
			newQ = append(newQ, [2]int{i - 1, j})
			newQ = append(newQ, [2]int{i, j + 1})
			newQ = append(newQ, [2]int{i, j - 1})
		}
		q = newQ
	}
	return -1
}
