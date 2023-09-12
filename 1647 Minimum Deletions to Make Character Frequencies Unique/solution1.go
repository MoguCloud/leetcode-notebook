// Your runtime beats 96.97 % of golang submissions (10 ms)
// Your memory usage beats 46.67 % of golang submissions (6.99 MB)
//
// 按字母出现频率从大到小排序
// 排在后面的要比排在前面的出现次数小
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func minDeletions(s string) int {
	counts := make([]int, 26)
	for _, str := range s {
		counts[str-'a']++
	}
	// 按字母出现频率从大到小排序
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	ret := 0
	i := 0
	j := 1
	for j < 26 && counts[j] != 0 {
		// 排在后面的要比排在前面的出现次数小
		if counts[i] <= counts[j] {
			// 要删除到比前面的小1
			count := counts[j] - counts[i] + 1
			ret += count
			counts[j] -= count
		}
		// 避免 [1,1,1] 删完变成 [1,0,1] 不进行后续比较
		// 将 counts[i] 和 counts[j] 进行互换
		// 将 [1, 0, 1] 变成 [0, 1, 1]
		if counts[j] == 0 {
			counts[i], counts[j] = counts[j], counts[i]
		}
		i = j
		j++
	}
	return ret
}
