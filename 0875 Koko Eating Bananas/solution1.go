// Your runtime beats 100 % of golang submissions (16 ms)
// Your memory usage beats 35.56 % of golang submissions (6.6 MB)
//
// 二分查找

func minEatingSpeed(piles []int, h int) int {
	lo := 1
	hi := 1000000000
	for lo < hi {
		mid := lo + (hi-lo)/2
		if canFinish(piles, h, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func canFinish(piles []int, h int, k int) bool {
	count := 0
	for _, p := range piles {
		count += int(math.Ceil(float64(p) / float64(k)))
		if count > h {
			return false
		}
	}
	return true
}
