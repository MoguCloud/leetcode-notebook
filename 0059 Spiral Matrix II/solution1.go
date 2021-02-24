// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 15.24 % of golang submissions (2.1 MB)
//
// 模拟螺旋走一遍，初始方向朝右。当下一格溢出或者(ret[x][y])值不为0时，则需要顺时针选择90度。

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	direction := 0 // 右 -> 下 -> 左 -> 上
	x := 0 // 当前点的座标 (x, y)
	y := 0
	for i := 0; i < n*n; i++ {
		ret[x][y] = i + 1
		switch direction % 4 {
		case 0:
			if y < n-1 && ret[x][y+1] == 0 {
				y += 1
			} else {
				direction += 1
				x += 1
			}
		case 1:
			if x < n-1 && ret[x+1][y] == 0 {
				x += 1
			} else {
				direction += 1
				y -= 1
			}
		case 2:
			if y > 0 && ret[x][y-1] == 0 {
				y -= 1
			} else {
				direction += 1
				x -= 1
			}
		case 3:
			if x > 0 && ret[x-1][y] == 0 {
				x -= 1
			} else {
				direction += 1
				y += 1
			}
		}
	}
	return ret
}
