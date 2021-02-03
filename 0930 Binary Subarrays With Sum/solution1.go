// Prefix Sum
//
// Your runtime beats 66.67 % of golang submissions (36 ms)
// Your memory usage beats 77.78 % of golang submissions (7.2 MB)
//
// 哈希表代替数组，哈希表key为前缀和，值为前缀和出现的次数
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func numSubarraysWithSum(A []int, S int) int {
	ret := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		count, ok := prefixSum[sum-S] // prefixSum[i] - prefixSum[j] == S
		if ok {
			ret += count
		}
		prefixSum[sum] += 1
	}
	return ret
}
