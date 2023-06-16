// Your runtime beats 35.71 % of golang submissions (50 ms)
// Your memory usage beats 42.86 % of golang submissions (7.2 MB)
//
// DP
// dp[move][num] 跳了move次跳到num构成的电话号码个数
// 初始状态 dp[0][num] = 1
// 状态转移方程 dp[move][num] = dp[move-1][num2] 之和，其中 num2 为 可以跳到 num 的数
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

var next map[int][]int = map[int][]int{
	1: []int{8, 6},
	2: []int{7, 9},
	3: []int{4, 8},
	4: []int{3, 9, 0},
	5: []int{},
	6: []int{7, 1, 0},
	7: []int{2, 6},
	8: []int{1, 3},
	9: []int{2, 4},
	0: []int{4, 6},
}

// dp[move][num] 跳了move次跳到num构成的电话号码个数
func knightDialer(n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 10)
	}
	for i := 0; i < 10; i++ {
		dp[0][i] = 1
	}
	for move := 1; move < n; move++ {
		for num := 0; num < 10; num++ {
			for _, dest := range next[num] {
				dp[move][dest] += dp[move-1][num]
				dp[move][dest] %= 1e9 + 7
			}
		}
	}
	ret := 0
	for num := 0; num < 10; num++ {
		ret += dp[len(dp)-1][num]
		ret %= 1e9 + 7
	}
	return ret
}
