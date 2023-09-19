// Your runtime beats 31.54 % of golang submissions (110 ms)
// Your memory usage beats 28.84 % of golang submissions (10.35 MB)
//
// 二分搜索
// counti为在数组中小于等于i出现的次数
// 如果数字j == countj，则重复数字 >  j
// 如果数字j >  countj，则重复数字 >  j
// 如果数字j <  countj，则重复数字 <= j
// 因为满足单调性所以可以使用二分搜索
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(1)

func findDuplicate(nums []int) int {
	left := 1
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
