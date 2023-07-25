// Your runtime beats 77.94 % of golang submissions (102 ms)
// Your memory usage beats 91.18 % of golang submissions (8.18 MB)
//
// 二分搜索
//
// 时间复杂度 O(logn)
// 空间复杂度 O(1)

func peakIndexInMountainArray(arr []int) int {
	lo := 0
	hi := len(arr) - 1
	for hi > lo {
		mid := lo + (hi-lo)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] < arr[mid+1] {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}
