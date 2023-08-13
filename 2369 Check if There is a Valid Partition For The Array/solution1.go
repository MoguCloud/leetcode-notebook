// Your runtime beats 42.86 % of golang submissions (198 ms)
// Your memory usage beats 85.71 % of golang submissions (9.16 MB)
//
// DP
// dp[i+1] 为 nums[0...i]是否能够被划分
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func validPartition(nums []int) bool {
	dp := make([]bool, len(nums)+1)
	dp[0] = true
	for i := 2; i < len(dp); i++ {
		if dp[i-2] && nums[i-1] == nums[i-2] {
			dp[i] = true
		} else if i >= 3 && dp[i-3] {
			if nums[i-1] == nums[i-2] && nums[i-2] == nums[i-3] {
				dp[i] = true
			} else if nums[i-1]-nums[i-2] == 1 && nums[i-2]-nums[i-3] == 1 {
				dp[i] = true
			}
		}
	}
	return dp[len(dp)-1]
}
