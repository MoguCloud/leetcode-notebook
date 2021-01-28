// Prefix Sum
//
// Your runtime beats 28.57 % of golang submissions (236 ms)
// Your memory usage beats 21.43 % of golang submissions (7.9 MB)
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)
//
// 0 <= j < i < len(nums)
// prefix[i] 为前i个数中奇数的个数
// prefix[i] = prefix[i - 1] + nums[i] & 1
// prefix[i] - prefix[j] 为 [j, i] 存在的奇数个数
// prefix[i] - prefix[j] == k
// 因为 prefix[i] 只与前一项有关系，所以可以用一个变量代替数组。
// 每次计算完 prefix[i] 后都需要在 prefix[0:i] 中找出 prefix[j] == prefix[i] - k 的个数，
//     可以用哈希表存放prefix[j]的个数。从 O(n) 变成 O(1)
func numberOfSubarrays(nums []int, k int) int {
	cnt := make(map[int]int) // key: prefix[j] val: prefix[j]的个数
	cnt[0] = 1
	sum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i] & 1
		cnt[sum] += 1
		if c, ok := cnt[sum-k]; ok {
			count += c
		}
	}
	return count
}
