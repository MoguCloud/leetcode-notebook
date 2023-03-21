// Your runtime beats 55.46 % of golang submissions (62 ms)
// Your memory usage beats 11.34 % of golang submissions (8.7 MB)
//
// BFS
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

func updateMatrix(mat [][]int) [][]int {
	c := len(mat) * len(mat[0])
	q := [][2]int{}
	// 使用mat作为结果数组
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				// 将所有未计算过距离的格子置为-1
				mat[i][j] = -1
			} else {
				// 从0的格子向外开始计算距离
				q = append(q, [2]int{i - 1, j})
				q = append(q, [2]int{i + 1, j})
				q = append(q, [2]int{i, j + 1})
				q = append(q, [2]int{i, j - 1})
				c--
			}
		}
	}
	distance := 1
	// 循环结束条件是所有格子的距离计算过1次
	// 因为是从所有的0开始向外bfs，所以能保证第一次计算到的距离就是最小值
	for c > 0 {
		newQ := [][2]int{}
		for i := 0; i < len(q); i++ {
			x := q[i][0]
			y := q[i][1]
			// mat >= 0 则该格子已经计算过距离
			if x < 0 || x >= len(mat) || y < 0 || y >= len(mat[0]) || mat[x][y] >= 0 {
				continue
			}
			if mat[x][y] == -1 {
				c--
				mat[x][y] = distance
				newQ = append(newQ, [2]int{x - 1, y})
				newQ = append(newQ, [2]int{x + 1, y})
				newQ = append(newQ, [2]int{x, y - 1})
				newQ = append(newQ, [2]int{x, y + 1})
			}
		}
		distance++
		q = newQ
	}
	return mat
}
