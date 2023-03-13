// Your runtime beats 98.21 % of golang submissions (6 ms)
// Your memory usage beats 26.20 % of golang submissions (5.5 MB)
//
// 双指针
// 起始时候左、右指针分别指着数组的最左和最右的值
// 左右指针之和等于target时，则是解
// 和>target，则右指针往左移动一位
// 和<target，则左指针往右移动一位
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right = right - 1
		} else {
			left = left + 1
		}
	}
	return []int{}
}
