// Your runtime beats 100 % of golang submissions (52 ms)
// Your memory usage beats 80 % of golang submissions (8.5 MB)
//
// maxSub 数组arr中的最大连续子序列和
// prefix 数组arr中的最大前缀和
// suffix 数组arr中的最大后缀和
//
// maxSub的计算方法可以通过 DP 求解
// dp[i] = max(dp[i-1], 0) + arr[i]
// maxSub = max(dp)
// 因为DP的状态转移方程只与前一项有关系，所以只需要一个变量就可以压缩空间求解
//
// 1. 当 k == 1 时
// 结果为 maxSub
//
// 2. 当 k > 1 时
//   I. 当 sum <= 0 时
//     结果为 max(maxSub, prefix + suffix)
//     因为 sum <= 0， 所以结果会出现在以下两种情况，所以取以下两种情况的最大值即可
//       i. 最大和出现在1个arr中，此时结果就为maxSub
//       ii. 最大和出现在2个arr中，即前一个arr最大后缀和拼接后一个arr的最大前缀和，此时结果为 suffix + prefix
//   II. 当 sum > 0 时
//     结果为 prefix + (k - 2) * sum + suffix
//     即 第一项取最大后缀和suffix 中间项取数组和sum 最后项取最大后缀和prefix
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func kConcatenationMaxSum(arr []int, k int) int {
	sum := 0
	maxSub := 0     // 最大子序列和
	prefix := 0     // 最大前缀和
	suffix := 0     // 最大后缀和
	prevMaxSub := 0 // 用于求最大子序列和
	for i := 0; i < len(arr); i++ {
		prevMaxSub = max(prevMaxSub, 0) + arr[i]
		if prevMaxSub > maxSub {
			maxSub = prevMaxSub
		}
		sum += arr[i]
		if sum > prefix {
			prefix = sum
		}
	}
	sum = 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i]
		if sum > suffix {
			suffix = sum
		}
	}

	if k == 1 {
		return maxSub
	} else {
		if sum > 0 {
			return int((int64(k-2)*int64(sum) + int64(prefix) + int64(suffix)) % 1000000007)
		} else {
			return max(prefix+suffix, maxSub)
		}
	}
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
