// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 48.62 % of golang submissions (2.53 MB)
//
// 二分搜索
// 先通过二分搜索找到旋转点
// 再判断target在旋转点之前还是旋转点之后
// 在通过二分搜索寻找target
//
// 时间复杂度 O(logn)
// 空间复杂度 O(1)

func search(nums []int, target int) int {
	n := 0
	var left, right, ret int
	if nums[len(nums)-1] < nums[0] {
		// 寻找旋转点
		left = 1
		right = len(nums) - 1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] < nums[mid-1] {
				n = mid
				break
			}
			if nums[mid] > nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	// 判断是在旋转点之前还是在旋转点之后，然后二分搜索
	if nums[n] <= target && target <= nums[len(nums)-1] {
		ret = binarySearch(nums, target, n, len(nums)-1)
	} else {
		ret = binarySearch(nums, target, 0, n-1)
	}
	if nums[ret] != target {
		return -1
	}
	return ret
}

func binarySearch(nums []int, target int, left int, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}
