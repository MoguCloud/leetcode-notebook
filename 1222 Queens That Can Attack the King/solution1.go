// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 75 % of golang submissions (2.6 MB)
//
// 从国王的位置从8个方向开始遍历，直到遇到皇后或者到边界

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	ret := [][]int{}
	chessboard := make([][]int, 8)
	for i := 0; i < 8; i++ {
		chessboard[i] = make([]int, 8)
	}
	for _, queen := range queens {
		chessboard[queen[0]][queen[1]] = 1
	}
	directions := [][]int{
		[]int{0, 1},
		[]int{0, -1},
		[]int{1, 0},
		[]int{-1, 0},
		[]int{1, 1},
		[]int{1, -1},
		[]int{-1, 1},
		[]int{-1, -1},
	}
	for _, d := range directions {
		k := []int{king[0], king[1]}
		for {
			if k[0] < 0 || k[0] >= 8 || k[1] < 0 || k[1] >= 8 {
				break
			}
			if chessboard[k[0]][k[1]] == 1 {
				ret = append(ret, k)
				break
			}
			k[0] += d[0]
			k[1] += d[1]
		}
	}
	return ret
}
