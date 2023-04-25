// Your runtime beats 60 % of golang submissions (189 ms)
// Your memory usage beats 40 % of golang submissions (9.9 MB)
//
// 贪心
// 回家的路线包括水平方向移动和垂直方向移动这2种移动方向，无论先水平方向移动还是先垂直方向移动还是交替移动，都不会影响最终结果，所以只要计算水平方向移动的开销和垂直方向移动的开销之和即可。
//
// 时间复杂度 O(m+n)
// 空间复杂度 O(1)

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	ret := 0
	if startPos[0] < homePos[0] {
		for i := startPos[0] + 1; i <= homePos[0]; i++ {
			ret += rowCosts[i]
		}
	}
	if startPos[0] > homePos[0] {
		for i := startPos[0] - 1; i >= homePos[0]; i-- {
			ret += rowCosts[i]
		}
	}
	if startPos[1] < homePos[1] {
		for i := startPos[1] + 1; i <= homePos[1]; i++ {
			ret += colCosts[i]
		}
	}
	if startPos[1] > homePos[1] {
		for i := startPos[1] - 1; i >= homePos[1]; i-- {
			ret += colCosts[i]
		}
	}
	return ret
}
