// Your runtime beats 96.97 % of golang submissions (35 ms)
// Your memory usage beats 63.64 % of golang submissions (7.32 MB)
//
// 枚举
// 两个相连的城市的秩   = 城市A的度 + 城市B的度 - 1
// 两个不相连的城市的秩 = 城市A的度 + 城市B的度 - 1
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func maximalNetworkRank(n int, roads [][]int) int {
	// 记录度
	degrees := make([]int, n)
	// 记录邻接关系
	graph := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]struct{})
	}
	for _, road := range roads {
		degrees[road[0]]++
		degrees[road[1]]++
		graph[road[0]][road[1]] = struct{}{}
		graph[road[1]][road[0]] = struct{}{}
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := degrees[i] + degrees[j]
			// 判断是否城市是否连接
			if _, ok := graph[i][j]; ok {
				rank--
			}
			if rank > max {
				max = rank
			}
		}
	}
	return max
}
