// Your runtime beats 64.71 % of golang submissions (8 ms)
// Your memory usage beats 52.94 % of golang submissions (7.2 MB)
//
// DP
// dp[i][j] 为 s1[0:i] 与 s2[0:j] 构成相同子串需要删除的最小ascii和
// 初始条件 dp[0][0] = 0, s1[0:0] 与 s2[0:0] 都是空字符串
//          dp[0][j] = dp[0][j-1] + s2[j-1], 将s2[0:j]变得和空字符串相同，需要删除s2中所有字符
//          dp[i][0] = dp[i-1][0] + s1[i-1], 将s1[0:i]变得和空字符串相同，需要删除s1中所有字符
// 状态转移方程 当 s1[i] == s2[j]时，dp[i][j] = dp[i-1][j-1]，即字符相同时不需要删除
//              当 s1[i] != s2[j]时，dp[i][j] = min(dp[i-1][j] + s1[i-1], dp[i][j-1] + s2[j-1])，即此时有2种可以选择的方案，选其中和为最小的情况
//                  1. 选择s1[0:i-1]和s2[0:j]构成的子串，删除s1[i-1]，此时的ascii删除和为 dp[i-1][j] + s1[i-1]
//                  2. 选择s1[0:i]和s2[0:j-1]构成的子串，删除s2[j-1]，此时的ascii删除和为 dp[i][j-1] + s2[j-1]
//
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
