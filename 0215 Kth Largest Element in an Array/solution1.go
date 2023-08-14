// Your runtime beats 75.29 % of golang submissions (83 ms)
// Your memory usage beats 41.28 % of golang submissions (8.5 MB)
//
// 快速选择
//
// 时间复杂度 O(n)
// 空间复杂度 O(logn)

func findKthLargest(nums []int, k int) int {
	// 第k大转换成第len(nums)-k小
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l int, r int, k int) int {
	p := partition(nums, l, r)
	if p == k {
		return nums[p]
	} else if p > k {
		return quickSelect(nums, l, p-1, k)
	} else {
		return quickSelect(nums, p+1, r, k)
	}
}

func partition(nums []int, l int, r int) int {
	pivot := nums[l]
	// nums在区间 [l+1, pivotIndex] 里的每个数都 <= pivot
	// nums在区间 (pivotIndex, r] 里的每个数都 > pivot
	// 初始情况下区间 [l+1, pivotIndex] 是空区间
	// 即 pivotIndex < l + 1   =>  pivotIndex = l
	pivotIndex := l
	for i := l + 1; i <= r; i++ {
		if nums[i] <= pivot {
			pivotIndex++
			nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
		}
	}
	// 将 pivot 与 nums[pivotIndex] 进行交换
	nums[l], nums[pivotIndex] = nums[pivotIndex], nums[l]
	return pivotIndex
}
