// Your runtime beats 80 % of golang submissions (156 ms)
// Your memory usage beats 20 % of golang submissions (10.7 MB)
//
// 排序
// 特殊楼层之间的连续楼层数 = special[i+1] - special[i] - 1
// 特殊楼层最低层和bottom之间的连续楼层数 = special[0] - bottom
// 特殊楼层最高层和top之间的连续楼层数 = top - special[n]
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(logn)

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	ret := 0
	for i := 1; i < len(special); i++ {
		ret = max(ret, special[i]-special[i-1]-1)
	}
	ret = max(ret, special[0]-bottom)
	ret = max(ret, top-special[len(special)-1])
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
