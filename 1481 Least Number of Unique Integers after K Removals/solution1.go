// Your runtime beats 9.09 % of golang submissions (428 ms)
// Your memory usage beats 81.82 % of golang submissions (11.4 MB)
//
// 优先将出现次数少的元素进行删除。
// 将数组按照出现次数进行排序，然后按照出现次数进行删除。
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	if k == len(arr) {
		return 0
	}
	intCount := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		intCount[arr[i]] += 1
	}
	// 按出现次数升序排序
	sort.Slice(arr, func(i, j int) bool {
		if intCount[arr[i]] < intCount[arr[j]] {
			return true
		} else if intCount[arr[i]] > intCount[arr[j]] {
			return false
		} else {
			return arr[i] < arr[j]
		}
	})
	for i := 0; i < k; i++ {
		intCount[arr[i]] -= 1
		if intCount[arr[i]] == 0 {
			delete(intCount, arr[i])
		}
	}
	return len(intCount)
}
