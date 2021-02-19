// Your runtime beats 63.64 % of golang submissions (108 ms)
// Your memory usage beats 77.27 % of golang submissions (11.5 MB)
func findLeastNumOfUniqueInts(arr []int, k int) int {
	if k == len(arr) {
		return 0
	}
	// 各元素出现的次数
	intCount := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		intCount[arr[i]] += 1
	}
	// arr中的元素去重然后按照出现次数排序
	countArr := []int{}
	for k, _ := range intCount {
		countArr = append(countArr, k)
	}
	sort.Slice(countArr, func(i, j int) bool {
		return intCount[countArr[i]] < intCount[countArr[j]]
	})
	ret := len(countArr)
	offset := 0
	for k > 0 {
		c := intCount[countArr[offset]]
		if k-c >= 0 {
			offset += 1
			ret -= 1
		}
		k = k - c
	}
	return ret
}
