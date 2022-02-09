// Your runtime beats 13.33 % of golang submissions (128 ms)
// Your memory usage beats 8.89 % of golang submissions (9.7 MB)
//
// DP
//
// minDp[i] 从 0 到 i 最小连续和
// maxDp[i] 从 0 到 i 最大连续和
// minDp 递推公式 minDp[i] = min(minDp[i-1] + nums[i], nums[i])
// maxDp 递推公式 maxDp[i] = max(maxDp[i-1] + nums[i], nums[i])
//
// 返回结果即 max(maxDp中最大值, minDp中最小值的绝对值)
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func maxAbsoluteSum(nums []int) int {
	maxDp := make([]int, len(nums))
	minDp := make([]int, len(nums))
	maxDp[0] = nums[0]
	minDp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		minDp[i] = min(minDp[i-1]+nums[i], nums[i])
		maxDp[i] = max(maxDp[i-1]+nums[i], nums[i])
	}
	maxSum := 0
	minSum := 0
	for i := 0; i < len(nums); i++ {
		if minDp[i] < minSum {
			minSum = minDp[i]
		}
		if maxDp[i] > maxSum {
			maxSum = maxDp[i]
		}
	}
	if -minSum > maxSum {
		return -minSum
	} else {
		return maxSum
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
