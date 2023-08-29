// Your runtime beats 66.67 % of golang submissions (6 ms)
// Your memory usage beats 100 % of golang submissions (5.57 MB)
//
// 计数
// 第i小时是否开店取决于后续Y和N的数量，如果Y的数量>N的数量，则应该开店
// 最优情况是 Y - N 的数量达到最大值时关门
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func bestClosingTime(customers string) int {
	maxP := 0
	p := 0
	ret := 0
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			p++
		} else {
			p--
		}
		if p > maxP {
			maxP = p
			ret = i + 1
		}
	}
	return ret
}
