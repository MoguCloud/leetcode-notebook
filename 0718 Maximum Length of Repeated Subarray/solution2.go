// Your runtime beats 55.7 % of golang submissions (72 ms)
// Your memory usage beats 35.44 % of golang submissions (16.6 MB)
//
// DP
//
// dp[i][j] 表示 数组nums1索引从i开始 数组nums2索引从j开始的最长公共前缀
// 状态转移方程:
// dp[i][j] = dp[i+1][j+1] + 1  (nums1[i] == nums2[j]时)
// dp[i][j] = 0                 (nums1[i] != nums2[j]时)
// 因为dp[i][j]的值依赖dp[i+1][j+1]，所以要倒序遍历
// 初始条件:
// dp[i][len(nums2)-1] = 1   (nums1[i] == nums2[len(nums2)-1]时)
// dp[i][len(nums2)-1] = 0   (nums1[i] != nums2[len(nums2)-1]时)
// 因为初始条件和状态转移方程一致，所以将dp数组大小定义成[len(nums1)+1][len(nums2)+1]即可
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n^2)
func findLength(nums1 []int, nums2 []int) int {
	maxLength := 0
	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(nums1)+1; i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return maxLength
}
