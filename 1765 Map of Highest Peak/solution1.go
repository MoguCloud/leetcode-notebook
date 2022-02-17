// Your runtime beats 69.57 % of golang submissions (352 ms)
// Your memory usage beats 26.09 % of golang submissions (45.2 MB)
//
// BFS

func highestPeak(isWater [][]int) [][]int {
	ret := make([][]int, len(isWater))
	queue := [][]int{}
	count := 0
	directions := [][]int{
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, 0},
		[]int{-1, 0},
	}
	// 进行初始化
	// 水高度置为0, 其他的高度设置为-1
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, len(isWater[i]))
		for j := 0; j < len(ret[i]); j++ {
			if isWater[i][j] == 1 {
				ret[i][j] = 0
				queue = append(queue, []int{i, j})
				count += 1
			} else {
				ret[i][j] = -1
			}
		}
	}
	for count < len(isWater)*len(isWater[0]) {
		p := queue[0]
		queue = queue[1:]
		for _, d := range directions {
			if p[0]+d[0] >= 0 && p[0]+d[0] < len(isWater) && p[1]+d[1] >= 0 && p[1]+d[1] < len(isWater[0]) {
				if ret[p[0]+d[0]][p[1]+d[1]] == -1 {
					ret[p[0]+d[0]][p[1]+d[1]] = ret[p[0]][p[1]] + 1
					count += 1
					queue = append(queue, []int{p[0] + d[0], p[1] + d[1]})
				}
			}
		}
	}
	return ret
}
