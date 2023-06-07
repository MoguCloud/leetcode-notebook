// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 25.93 % of golang submissions (1.9 MB)
//
// 按位遍历分别判断每一位需要翻转的次数
//
// 时间复杂度 O(logn)
// 空间复杂度 O(1)

func minFlips(a int, b int, c int) int {
	count := 0
	maxNum := max(a, b, c)
	i := 0
	for maxNum > 0 {
		if c&(1<<i) == 0 {
			if a&(1<<i) != 0 {
				count++
			}
			if b&(1<<i) != 0 {
				count++
			}
		} else {
			if a&(1<<i) == 0 && b&(1<<i) == 0 {
				count++
			}
		}
		i++
		maxNum >>= 1
	}
	return count
}

func max(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}
