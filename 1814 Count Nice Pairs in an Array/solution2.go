// Your runtime beats 20 % of golang submissions (104 ms)
// Your memory usage beats 20 % of golang submissions (9.4 MB)
//
// nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
// => nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
// 也就是任意两个 nums[i] - rev(nums[i]) 相同的数可以组成一对好对子
// 通过一个哈希表存放 nums[i] - rev(nums[i])
// 哈希表的 key 为 nums[i] - rev(nums[i])
// 哈希表的 val 为 nums[i] - rev(nums[i]) 的个数，记为n
// 那么可以组成 n * (n - 1) / 2 对好对子
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func countNicePairs(nums []int) int {
	var ret64 int64
	countMap := make(map[int]int) // key: nums[i] - rev(nums[i])  value: count
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]-rev(nums[i])] += 1
	}
	for _, v := range countMap {
		if v > 1 {
			ret64 += int64(v * (v - 1) / 2)
		}
	}
	return int(ret64 % 1000000007)
}

func rev(n int) int {
	ret := 0
	for n > 0 {
		ret = ret*10 + n%10
		n = n / 10
	}
	return ret
}
