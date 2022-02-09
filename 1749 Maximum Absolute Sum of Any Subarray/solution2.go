// Your runtime beats 24.44 % of golang submissions (114 ms)
// Your memory usage beats 24.44 % of golang submissions (9.1 MB)
//
// DP 空间压缩
//
// minDp[i] 从 0 到 i 最小连续和
// maxDp[i] 从 0 到 i 最大连续和
// minDp 递推公式 minDp[i] = min(minDp[i-1] + nums[i], nums[i])
// maxDp 递推公式 maxDp[i] = max(maxDp[i-1] + nums[i], nums[i])
//
// 返回结果即 max(maxDp中最大值, minDp中最小值的绝对值)
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func maxAbsoluteSum(nums []int) int {
	maxDp := nums[0]
	minDp := nums[0]
	maxSum := maxDp
	minSum := minDp
	for i := 1; i < len(nums); i++ {
		minDp = min(minDp+nums[i], nums[i])
		maxDp = max(maxDp+nums[i], nums[i])
		minSum = min(minSum, minDp)
		maxSum = max(maxSum, maxDp)
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
