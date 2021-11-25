// Your runtime beats 14.29 % of golang submissions (23 ms)
// Your memory usage beats 76.19 % of golang submissions (5.4 MB)
//
// 按区间的起始点升序排序，相同的则按结束点降序排序
// 顺序遍历数组，并且维护一个最大结束点 maxRight
// 如果当前点的结束点 < maxRight，则在当前点之前存在一个起始点<当前点并且结束点>当前点的，即存在一个区间可以覆盖当前区间

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		} else {
			return false
		}
	})
	count := len(intervals)
	maxRight := 0
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] > maxRight {
			maxRight = intervals[i][1]
		} else {
			count -= 1
		}
	}
	return count
}
