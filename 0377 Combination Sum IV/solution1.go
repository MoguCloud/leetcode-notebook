// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 34.18 % of golang submissions (2.00 MB)
//
// DP
// 类似爬楼梯
// dp[i] 为 可以组合成总数为 i 的个数
//
// 时间复杂度 O(target*len(nums)
// 空间复杂度 O(target)

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
