// Binary search
//
// Your runtime beats 98.77 % of golang submissions (4 ms)
// Your memory usage beats 20.99 % of golang submissions (4.1 MB)

func singleNonDuplicate(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := (hi + lo) / 2
		// 单一元素只会在奇数个的那边
		if nums[mid] == nums[mid-1] {
			// 如果 mid + 1 到 hi 为奇数个 则从 [mid + 1, hi] 中寻找，否则从[lo, mid] 中寻找
			if (hi-(mid+1)+1)&1 == 1 {
				lo = mid + 1
			} else {
				hi = mid
			}
		} else if nums[mid] == nums[mid+1] {
			// 如果 lo 到 mid - 1 为奇数个 则从 [lo, mid - 1] 中寻找，否则从[mid, hi] 中寻找
			if (mid-1-lo+1)&1 == 1 {
				hi = mid - 1
			} else {
				lo = mid
			}
		} else {
			return nums[mid]
		}
	}
	return nums[lo]
}
