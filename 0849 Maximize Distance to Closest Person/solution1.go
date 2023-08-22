// Your runtime beats 40.80 % of golang submissions (13 ms)
// Your memory usage beats 34.48 % of golang submissions (5.32 MB)
//
// 双指针
// 双指针分别指向前后2个有人的座位
// 如果第一个/最后一个座位没人，需要计算第一个/最后一个有人的座位到第一个/最后一个座位的距离
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func maxDistToClosest(seats []int) int {
	i := 0
	j := 0
	ret := 1
	for j < len(seats) {
		// 处理第一个座位是没人的情况
		if i == 0 && seats[0] == 0 && seats[j] == 1 {
			d := j - i
			ret = max(ret, d)
		}
		// 计算 2 个有人的座位之间的距离
		if seats[j] == 1 {
			d := (j - i) / 2
			ret = max(ret, d)
			i = j
		}
		j++
	}
	// 处理最后一个座位是没人的情况
	if i != len(seats)-1 {
		d := len(seats) - 1 - i
		ret = max(ret, d)
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
