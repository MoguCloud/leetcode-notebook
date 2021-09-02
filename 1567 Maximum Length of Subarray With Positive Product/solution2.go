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
// 因为 positive[n] 和 negative[n] 只与 positive[n-1] 和 negative[n-1] 有关
// 所以可以用 positive 和 negative 两个整型变量进行空间压缩
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func getMaxLen(nums []int) int {
	positive := 0
	negative := 0
	maxLen := 0
	if nums[0] > 0 {
		positive = 1
		maxLen = 1
	} else if nums[0] < 0 {
		negative = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			positive = positive + 1
			if negative != 0 {
				negative = negative + 1
			} else {
				negative = 0
			}
		} else if nums[i] < 0 {
			newNegative := positive + 1
			if negative != 0 {
				positive = negative + 1
			} else {
				positive = 0
			}
			negative = newNegative
		} else {
			positive = 0
			negative = 0
		}
		if positive > maxLen {
			maxLen = positive
		}
	}
	return maxLen
}
