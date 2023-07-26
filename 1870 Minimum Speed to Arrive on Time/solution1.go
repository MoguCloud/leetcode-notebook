// Your runtime beats 85.71 % of golang submissions (182 ms)
// Your memory usage beats 57.14 % of golang submissions (8.46 MB)
//
// 二分搜索
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(1)

func minSpeedOnTime(dist []int, hour float64) int {
	// 因为除了最后一段以外，每一段列车需要的最短时间是1
	// 最后一段需要的时间 > 0
	// 所以大于最短时间的直接返回 -1
	if int(math.Ceil(hour)) <= len(dist)-1 {
		return -1
	}
	// 速度的下界为 1
	// 上界为 10^7
	lo := 1
	hi := int(1e7)
	for lo < hi {
		mid := lo + (hi-lo)/2
		sum := 0.0
		// 除了最后一段以外，花费的时间需要向上取整
		for i := 0; i < len(dist)-1; i++ {
			sum += math.Ceil(float64(dist[i]) / float64(mid))
		}
		// 最后一段不需要向上取整
		sum += float64(dist[len(dist)-1]) / float64(mid)
		if sum > hour {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
