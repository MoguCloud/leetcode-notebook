// Your runtime beats 66.67 % of golang submissions (4 ms)
// Your memory usage beats 77.78 % of golang submissions (2.4 MB)
//
// 类选择排序
// 可以通过翻转2次将下标为k1的元素移动到下标为k2的位置
// 1. 翻转arr[0...k1]， 将元素移动到下标为0的位置
// 2. 翻转arr[0...k2]， 将元素从下标为0d的位置移动到下标为k2的位置
// 遍历n次（n = [len(arr)-1, 0] ），每次遍历将arr[0...n]中最大的元素移动到下标为n的位置即可
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)

func pancakeSort(arr []int) []int {
	ret := []int{}
	// 遍历 len(arr) - 1 次，每次将最大的元素移动到数组尾部
	for i := len(arr) - 1; i >= 0; i-- {
		maxIdx := i
		// 找出当前轮最大的元素
		for j := 0; j < i; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		// 如果最大元素已经在最后了，则不用翻转
		if maxIdx != i {
			// 翻转2次将最大的元素移动到数组尾部
			ret = append(ret, maxIdx+1)
			ret = append(ret, i+1)
			// 翻转数组
			for k := 0; k <= maxIdx/2; k++ {
				arr[k], arr[maxIdx-k] = arr[maxIdx-k], arr[k]
			}
			for k := 0; k <= i/2; k++ {
				arr[k], arr[i-k] = arr[i-k], arr[k]
			}
		}
	}
	return ret
}
