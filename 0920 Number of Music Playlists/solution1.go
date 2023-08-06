// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (4.23 MB)
//
// DP
//
// 时间复杂度 O(n * goal)
// 空间复杂度 O(n * goal)

// dp[i][j] 为 播放列表长度为i的，有j首不同的歌，满足条件的播放列表数量
// 问题的解就变成了求解 dp[goal][n]
func numMusicPlaylists(n int, goal int, k int) int {
	dp := make([][]int, goal+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1 // 初始状态 dp[0][0] = 1 初始状态只有一种选择的可能
	for i := 1; i <= goal; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] += dp[i-1][j] * max(j-k, 0)   //  第i首选择前i-1首歌里播放过的，并且要选择满足条件2的(一首歌只有在其他 k 首歌播放完之后才能再次播放。)，即j必须大于k，共有 j - k 首
			dp[i][j] += dp[i-1][j-1] * (n - j + 1) // 第i首选择前i-1首歌里没有播放过的，共有n-j+1首
			dp[i][j] %= 1000000007
		}
	}
	return dp[goal][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
