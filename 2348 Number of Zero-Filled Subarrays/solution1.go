// Your runtime beats 57.14 % of golang submissions (144 ms)
// Your memory usage beats 71.43 % of golang submissions (9.5 MB)
//
// 滑动窗口
// 找出所有连续的0，在计算0的子数组个数
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func zeroFilledSubarray(nums []int) int64 {
	// 方便处理最后一个数是0的情况
	nums = append(nums, -1)
	i := 0
	j := 0
	var ret int64 = 0
	for i <= len(nums)-1 {
		// 找出连续的0
		if nums[i] == 0 {
			j = i
			for nums[j] == 0 {
				j++
			}
			// 求连续0的子数组
			ret += int64(subLen(i, j))
			i = j
		} else {
			i++
		}
	}
	return ret
}

func subLen(i int, j int) int {
	// 长度为n的连续0的子数组
	// 1个长度为n 2个长度为n-1 ... n个长度为1
	// 子数组个数为 1 + 2 + ... + n
	length := j - i
	return (1 + length) * length / 2
}
