// Your runtime beats 46.97 % of golang submissions (43 ms)
// Your memory usage beats 56.44 % of golang submissions (7.3 MB)
//
// 双指针
// 当数组元素全为0时，长度为0
// 当数组元素全为1时，长度为数组长度-1
// 遍历数组，当数组元素为0时，删除该元素后子数组长度为该元素前面连续1的长度+该元素后面连续1的长度
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func longestSubarray(nums []int) int {
	// 判断数组是否包含1
	hasOne := false
	// 判断数组是否包含0
	hasZero := false
	// 前面连续1的长度
	prevOne := 0
	// 最长子数组长度
	max := 0
	i := 0
	for i < len(nums) {
		if nums[i] == 0 {
			hasZero = true
			// 当前元素为0，当前元素前一个元素也为0，则把前连续子数组长度置0
			if i > 0 && nums[i-1] == 0 {
				prevOne = 0
			}
			j := i + 1
			for j < len(nums) {
				if nums[j] == 1 {
					hasOne = true
					j++
				} else {
					break
				}
			}
			// j - i - 1 为后续连续子数组长度
			if j-i+prevOne > max {
				max = j - i - 1 + prevOne
			}
			prevOne = j - i - 1
			i = j
		} else {
			hasOne = true
			prevOne++
			i++
		}
	}
	// 当数组没有0，则数组全为1，子数组长度为数组长度-1
	if !hasZero {
		return len(nums) - 1
	}
	// 当数组没有1，则数组全为0，子数组长度为0
	if !hasOne {
		return 0
	}
	return max
}
