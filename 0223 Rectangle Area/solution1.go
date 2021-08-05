// Your runtime beats 28.57 % of golang submissions (16 ms)
// Your memory usage beats 9.52 % of golang submissions (6.3 MB)
//
// 两个矩形不重叠时，面积=矩形面积1+矩形面积2
//          重叠时，面积=矩形面积1+矩形面积2-重叠面积

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	s1 := (ax2 - ax1) * (ay2 - ay1)
	s2 := (bx2 - bx1) * (by2 - by1)
	cover := 0
	if ax2 < bx1 || bx2 < ax1 || ay2 < by1 || by2 < ay1 { // 判断是否重叠
		cover = 0
	} else {
		// 当线段有重叠时，重叠的长度=min(右值1, 右值2) - max(左值1, 左值2)
		cover = (min(ax2, bx2) - max(ax1, bx1)) * (min(ay2, by2) - max(ay1, by1))
	}
	return s1 + s2 - cover
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
