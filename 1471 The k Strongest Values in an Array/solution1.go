// Your runtime beats 50 % of golang submissions (248 ms)
// Your memory usage beats 50 % of golang submissions (8.9 MB)
//
// 自定义排序
//
// 时间复杂度 O(nlogn)

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	m := arr[(len(arr)-1)/2]

	sort.Slice(arr, func(i, j int) bool {
		if abs(arr[i]-m) > abs(arr[j]-m) {
			return true
		}
		if (abs(arr[i]-m) == abs(arr[j]-m)) && (arr[i] > arr[j]) {
			return true
		}
		return false
	})
	return arr[:k]
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
