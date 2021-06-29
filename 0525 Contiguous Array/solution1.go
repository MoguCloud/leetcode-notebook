// Your runtime beats 95.35 % of golang submissions (96 ms)
// Your memory usage beats 53.49 % of golang submissions (7.3 MB)
//
// 前缀和 + 哈希表
//
// 前缀和数组 prefix 遇到 0 则 -1，遇到 1 则 +1 (为了构建前缀和，需要满足遇到0和遇到1能够相互抵消)
// prefix[i] - prefix[j] == 0 时
//   则 数组 nums 在区间[j + 1, i] 内的 0 和 1 个数一样
//   即 子数组长度为 i - j
//
// 通过哈希表 memo 可以将当前索引i的前缀和记录在哈希表中，键为prefix[i]，值为i
// 遍历过程中看哈希表中是否存在键为prefix[i]的值
//   若存在则表示存在过一个j (i < j)，prefix[i] == prefix[j]
//   即 存在一个子数组 (nums[j + 1: i])内的 0 和 1 个数一样
//
// 因为使用了哈希表，prefix[i] 只与 prefix[i-1] 有关，并且 prefix 的历史信息存在哈希表内
// 所以 prefix 可以用单个变量单体数组，进行空间压缩
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func findMaxLength(nums []int) int {
	memo := make(map[int]int)
	memo[0] = -1
	prefixSum := 0
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			prefixSum -= 1
		} else {
			prefixSum += 1
		}
		if idx, ok := memo[prefixSum]; ok {
			length := i - idx
			if length > maxLen {
				maxLen = length
			}
		} else {
			memo[prefixSum] = i
		}
	}
	return maxLen
}
