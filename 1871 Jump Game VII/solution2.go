// Your runtime beats 86.67 % of golang submissions (12 ms)
// Your memory usage beats 66.67 % of golang submissions (7.5 MB)
//
// 动态规划 + 前缀和
//
// 如果 s[j] 可达，需要满足
// 1. s[j] == '0'
// 2. 存在任意一个s[i]可达， i 在 [j-maxJump, j-minJump]的范围内
//
// dp[i] 为 从起始位置能够跳到i
// dp[i] = any(dp[j]) j 在 [j-maxJump, j-minJump]的范围内
// 起始条件 dp[0] = true
//
// 可以通过前缀和在O(1)的时间内判断 是否存在任意一个s[i]可达， i 在 [j-maxJump, j-minJump]的范围内
// s[i] 可达 时记作1， s[i] 不可达 时记作0，只要判断前缀和数组的差是否 > 0 即可判断
// arr         1 1 0 1
// prefixSum   1 2 2 3
// sum(arr[i..j]) = prefixSum(j) - prefixSum(i-1)
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func canReach(s string, minJump int, maxJump int) bool {
	if s[len(s)-1] == '1' {
		return false
	}
	dp := make([]bool, len(s))
	prefixSum := make([]int, len(s))
	// 0 可达， (0, minJump) 不可达
	for i := 0; i < minJump; i++ {
		prefixSum[i] = 1
		dp[i] = false
	}
	dp[0] = true
	for i := minJump; i < len(s); i++ {
		canReach := false
		if s[i] == '0' {
			right := prefixSum[i-minJump]
			left := 0
			// 防止 i -maxJump - 1 < 0 时下标越界
			if i-maxJump-1 >= 0 {
				left = prefixSum[i-maxJump-1]
			}
			if right-left > 0 {
				canReach = true
			}
		}
		if canReach {
			dp[i] = true
			prefixSum[i] = prefixSum[i-1] + 1
		} else {
			dp[i] = false
			prefixSum[i] = prefixSum[i-1]
		}
	}
	return dp[len(dp)-1]
}
