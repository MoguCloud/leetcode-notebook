// Your runtime beats 85.71 % of golang submissions (88 ms)
// Your memory usage beats 33.33 % of golang submissions (10.3 MB)
//
// DP
//
// positive[i] 为 下标为 i 的乘积为正数的最长数组长度
// negative[i] 为 下标为 i 的乘积为负数的最长数组长度
//
// 状态转移方程
// 当 nums[0] > 0 时，
//     positive[i] = positive[i-1] + 1
//   当 negative[i-1] == 0 时，
//     negative[i] = 0
//   当 negative[i-1] > 0 时，
//     negative[i] = negative[i-1] + 1
// 当 nums[0] < 0 时，
//     positive[i] = negative[i-1] + 1
//   当 negative[i-1] == 0 时，
//     positive[i] = 0
//   当 negative[i-1] > 0 时，
//     positive[i] = negative[i-1] + 1
// 初始条件
// 当 nums[0] > 0 时，positive[i] = 1
// 当 nums[0] < 0 时，negative[i] = 1
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func getMaxLen(nums []int) int {
	positive := make([]int, len(nums))
	negative := make([]int, len(nums))
	maxLen := 0
	if nums[0] > 0 {
		positive[0] = 1
		maxLen = 1
	} else if nums[0] < 0 {
		negative[0] = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			positive[i] = positive[i-1] + 1
			if negative[i-1] != 0 {
				negative[i] = negative[i-1] + 1
			}
		} else if nums[i] < 0 {
			negative[i] = positive[i-1] + 1
			if negative[i-1] != 0 {
				positive[i] = negative[i-1] + 1
			}
		} else {
			positive[i] = 0
			negative[i] = 0
		}
		if positive[i] > maxLen {
			maxLen = positive[i]
		}
	}
	return maxLen
}
