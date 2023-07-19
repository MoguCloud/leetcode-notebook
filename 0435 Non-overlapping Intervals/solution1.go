// Your runtime beats 26.55 % of golang submissions (262 ms)
// Your memory usage beats 11.86 % of golang submissions (16.93 MB)
//
// 贪心
// 按结束从小到大排序，每次都选择最先结束的
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(logn)

func eraseOverlapIntervals(intervals [][]int) int {
	// 按结束从小到大排序，结束相同的按开始从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ret := 0
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		// 结果和上一个一样，则会重叠
		if intervals[i][1] == prev[1] {
			ret++
		} else {
			// 结束比上一个大的时，再判断开始是否大于上一个结束
			if intervals[i][0] >= prev[1] {
				// 如果大于等于上一个的话，则不会覆盖，这个区域可以被选择
				// 下一次循环和这个区间进行比较
				prev = intervals[i]
			} else {
				ret++
			}
		}
	}
	return ret
}
