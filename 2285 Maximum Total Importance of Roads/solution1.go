// Your runtime beats 91.67 % of golang submissions (219 ms)
// Your memory usage beats 50 % of golang submissions (14.2 MB)
//
// 贪心
// 将道路按连接城市的数量进行排序，道路按连接数量从大到小给予从大到小的重要性，连接数量最多的给最大重要性，连接数量最少的给最小重要性
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

func maximumImportance(n int, roads [][]int) int64 {
	conns := make([]int, n)
	for _, road := range roads {
		conns[road[0]]++
		conns[road[1]]++
	}
	sort.Ints(conns)
	var ret int64
	for n > 0 {
		ret += int64(n * conns[n-1])
		n--
	}
	return ret
}
