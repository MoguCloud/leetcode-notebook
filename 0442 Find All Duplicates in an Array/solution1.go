// 原地 hash
// Your runtime beats 99.36 % of golang submissions (44 ms)
// Your memory usage beats 94.27 % of golang submissions (6.9 MB)
//
// 因为 nums 的值域在 [1, len(nums)]
// 所以可以用 nums[nums[i] - 1] 对 nums[i] 状态进行标记
// 当 nums[i] 的值出现过，将 nums[nums[i] - 1] 变成负的 nums[nums[i] - 1]
// 所以当 nums[nums[i] - 1] 为负数时，则 nums[i] 已经出现过

func findDuplicates(nums []int) []int {
	duplicatedArr := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[abs(nums[i])-1] < 0 {
			duplicatedArr = append(duplicatedArr, abs(nums[i]))
		} else {
			nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		}
	}
	return duplicatedArr
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
