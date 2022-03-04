// Your runtime beats 64.71 % of golang submissions (192 ms)
// Your memory usage beats 58.82 % of golang submissions (10.7 MB)
//
// 计算怪物抵达的时间，并且进行排序
// 从抵达时间由近到远进行攻击
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

func eliminateMaximum(dist []int, speed []int) int {
	time := make([]float64, len(dist))
	for i := 0; i < len(dist); i++ {
		time[i] = float64(dist[i]) / float64(speed[i])
	}
	sort.Float64s(time)
	ret := 0
	for i := 0; i < len(time); i++ {
		if time[i] > float64(i) {
			ret += 1
		} else {
			break
		}
	}
	return ret
}
