// Your runtime beats 100.00 % of golang submissions (4 ms)
// Your memory usage beats 66.67 % of golang submissions (2.3 MB)
//
// 状态压缩动态规划
// dp[mask]，（其中i=二进制mask中1的位数） 代表mask为1的位数上的导师分配了i个学生的兼容性评分和
// dp[i] = max{dp[mask']+cap[j][i]} 其中j为mask中为1的位数，mask'为将mask的第j位变为0
//
// 时间复杂度 O(m*m*n+m*2^m)
// 空间复杂度 O(m*m+2^m)

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	cap := make([][]int, m)
	dp := make([]int, 1<<m)
	for i := 0; i < m; i++ {
		cap[i] = make([]int, m)
	}
	// 计算所有学生和导师的兼容性评分
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			cap[i][j] = getCompatibility(students[i], mentors[j])
		}
	}
	for i := 1; i < (1 << m); i++ {
		for j := 0; j < m; j++ {
			if i&(1<<j) > 0 {
				dp[i] = max(dp[i], dp[i^(1<<j)]+cap[countOnes(i)-1][j])
			}
		}
	}
	return dp[len(dp)-1]
}

func getCompatibility(s []int, m []int) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		if s[i] == m[i] {
			ret += 1
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// countOnes 计算二进制中1的数量
func countOnes(n int) int {
	ret := 0
	for n > 0 {
		if n&1 == 1 {
			ret += 1
		}
		n = n >> 1
	}
	return ret
}
