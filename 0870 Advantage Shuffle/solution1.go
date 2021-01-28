// 贪心
//
// Your runtime beats 70 % of golang submissions （92 ms)
// Your memory usage beats 20 % of golang submissions (7.6 MB)
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)
//
// 分别将 A 和 B 排序，如果 A 中的牌 a 大于 B 中的牌 b，则将 a 放入 b 原来的位置；
//     若不大于，则将 a 放到 B 中最大的牌 b 原来的位置。
// 因为 B 中的牌会有重复，所以哈希表值的类型是数组。

func advantageCount(A []int, B []int) []int {
	bPositionMap := make(map[int][]int)
	for i := 0; i < len(A); i++ {
		arr, ok := bPositionMap[B[i]]
		if !ok {
			arr = []int{}
		}
		arr = append(arr, i)
		bPositionMap[B[i]] = arr
	}
	sort.Ints(A)
	sort.Ints(B)
	ret := make([]int, len(A))
	endB := len(B) - 1
	startB := 0
	for i := 0; i < len(A); i++ {
		if A[i] > B[startB] {
			arr, _ := bPositionMap[B[startB]]
			ret[arr[len(arr)-1]] = A[i]
			bPositionMap[B[startB]] = arr[:len(arr)-1]
			startB += 1
		} else {
			arr, _ := bPositionMap[B[endB]]
			ret[arr[len(arr)-1]] = A[i]
			bPositionMap[B[endB]] = arr[:len(arr)-1]
			endB -= 1
		}
	}
	return ret
}
