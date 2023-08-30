// Your runtime beats 100 % of golang submissions (84 ms)
// Your memory usage beats 50 % of golang submissions (9.74 MB)
//
// 贪心
// 从后往前遍历，如果当前元素小于等于后一个元素，则不需要拆分
// 当拆分的时候需要保证最大值小于等于后一个元素，最小值要尽可能大
// 就需要尽可能平均拆分
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func minimumReplacement(nums []int) int64 {
	// 用于维护上一个元素或者拆分完的最小值
	prev := nums[len(nums)-1]
	var ret int64
	for i := len(nums) - 2; i >= 0; i-- {
		// 不需要拆分
		if nums[i] <= prev {
			// 更新prev的值
			prev = nums[i]
			continue
		}
		// 对nums[i]进行k等分，保证最大值小于等于 nums[i] / prev
		// 等价于 ceil( nums[i] / prev )
		k := (nums[i] + prev - 1) / prev
		// 最小值为 nums[i] / k 下取整
		prev = nums[i] / k
		// 进行 k - 1 次操作即可 k 等分
		ret += int64(k - 1)
	}
	return ret
}
