// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 98.33 % of golang submissions (2.66 MB)
//
// 二分搜索
// 首先使用二分搜索来确定target在matrix的哪一行
// 再使用二分搜索来确定target是否在那一行里
//
// 时间复杂度 O(logm + logn)
// 空间复杂度 O(1)

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}
	low := 0
	high := m - 1
	for low <= high {
		mid := low + (high-low)/2
		if target >= matrix[mid][0] && target <= matrix[mid][n-1] {
			// 确定完是哪一行了，再去那一行里面搜索
			return searchArray(matrix[mid], target)
		}
		if target < matrix[mid][0] {
			high = mid - 1
		}
		if target > matrix[mid][n-1] {
			low = mid + 1
		}
	}
	return false
}

func searchArray(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if target == arr[mid] {
			return true
		}
		if target > arr[mid] {
			low = mid + 1
		}
		if target < arr[mid] {
			high = mid - 1
		}
	}
	return false
}
