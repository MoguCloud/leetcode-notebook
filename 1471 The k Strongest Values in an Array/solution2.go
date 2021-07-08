// Your runtime beats 75 % of golang submissions (200 ms)
// Your memory usage beats 25 % of golang submissions (9 MB)
//
// 双指针
//
// 先排序求出中位数
// 根据题意，比较的大小其实是到中位数的距离(和中位数差的绝对值)
// 也就是取k个到中位数距离最大的数
// 因为到中位数距离最大的只会是排序后最左或者最右的元素
// 所以可以靠双指针进行求解
//
// 时间复杂度 O(nlogn)

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	m := arr[(len(arr)-1)/2]
	ret := []int{}

	for i, j := 0, len(arr)-1; len(ret) < k; {
		if m-arr[i] > arr[j]-m {
			ret = append(ret, arr[i])
			i += 1
		} else {
			ret = append(ret, arr[j])
			j -= 1
		}
	}
	return ret
}
