// Your runtime beats 62.89 % of golang submissions (2 ms)
// Your memory usage beats 85.57 % of golang submissions (2.1 MB)
//
// 滑动窗口
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func getAverages(nums []int, k int) []int {
	ret := make([]int, len(nums))
	// 处理左边元素不足的情况
	for i := 0; i < k && i < len(nums); i++ {
		ret[i] = -1
	}
	// 处理右边元素不足的情况
	for i := len(nums) - k; i < len(nums) && i >= 0; i++ {
		ret[i] = -1
	}
	// 处理中间的平均值
	if k < len(nums)-k {
		sum := 0
		for i := 0; i < 2*k+1; i++ {
			sum += nums[i]
		}
		l := 0
		r := 2*k + 1
		for i := k; i < len(nums)-k; i++ {
			ret[i] = sum / (2*k + 1)
			// 防止下标越界
			if r < len(nums) {
				sum = sum - nums[l] + nums[r]
				l++
				r++
			} else {
				break
			}
		}
	}
	return ret
}
