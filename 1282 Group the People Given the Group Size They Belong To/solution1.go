// Your runtime beats 71.43 % of golang submissions (7 ms)
// Your memory usage beats 35.71 % of golang submissions (5.52 MB)
//
// Hash Table
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func groupThePeople(groupSizes []int) [][]int {
	groupMap := make(map[int][]int)
	for i, size := range groupSizes {
		groupMap[size] = append(groupMap[size], i)
	}
	ret := [][]int{}
	for size, groups := range groupMap {
		for i := 0; i < len(groups); i += size {
			group := groups[i : i+size]
			ret = append(ret, group)
		}
	}
	return ret
}
