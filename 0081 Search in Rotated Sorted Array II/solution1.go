// Your runtime beats 46.86 % of golang submissions (6 ms)
// Your memory usage beats 98.86 % of golang submissions (3.2 MB)
//
// 二分搜索
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)


func search(nums []int, target int) bool {
	pivot := 0
	// 处理类型[1,1,1,1,0,1]这种无法通过二分来寻找旋转点的情况
	for len(nums) > 1 && nums[0] == nums[len(nums)-1] {
		if nums[0] == target {
			return true
		}
		nums = nums[1 : len(nums)-1]
	}
	// 通过二分来寻找旋转点
	left := 1
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid-1] > nums[mid] {
			pivot = mid
			break
		}
		if nums[mid] < nums[0] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 根据旋转点和target的大小关系进行部分二分搜索
	if pivot == 0 {
		return binarySearch(nums, target)
	} else if target >= nums[0] && target <= nums[pivot-1] {
		return binarySearch(nums[:pivot], target)
	} else {
		return binarySearch(nums[pivot:len(nums)], target)
	}
}

func binarySearch(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
