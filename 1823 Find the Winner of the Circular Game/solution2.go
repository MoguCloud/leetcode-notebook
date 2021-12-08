// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (2 MB)
//
// DP
// 
// 假设 y = f(n, k) 返回胜利者
// 最开始状态 1, 2, ..., n - 1, n
// 淘汰了第k个 1, 2, ..., k - 1, k + 1, ...,  n - 1, n
// 下一轮状态 k + 1, k + 2, ..., n - 1, n, 1, 2, ..., k - 1
//
// 第1轮 1, 2, ..., k, ..., y, .. n - 1, n
// 第2轮 k + 1, ... y, ... n, ..., k - 1
// f(n, k) = (f(n-1, k) + k) % n
// f(1, k) = 1

func findTheWinner(n int, k int) int {
	dp := 0
	for i := 2; i <= n; i++ {
		dp = (dp + k) % i
	}
	return dp + 1
}
