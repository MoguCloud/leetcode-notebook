// Your runtime beats 94.51 % of golang submissions (24 ms)
// Your memory usage beats 90.30 % of golang submissions (6.9 MB)
//
// 贪心
// 将数组对按right降序排序
// 每次都取满足条件的最小right即可找到最长数对链
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(1)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	ret := 1
	prev := pairs[0][1]
	for _, pair := range pairs[1:len(pairs)] {
		if pair[0] > prev {
			ret++
			prev = pair[1]
		}
	}
	return ret
}
