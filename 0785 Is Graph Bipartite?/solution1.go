// Your runtime beats 98.67 % of golang submissions (15 ms)
// Your memory usage beats 34.67 % of golang submissions (6.92 MB)
//
// DFS
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

// visted 数组用于记录节点是否被标记过，0为标记，1为标记成集合A，2为标记成集合B
var visited []int
var result bool

func isBipartite(graph [][]int) bool {
	visited = make([]int, len(graph))
	result = true
	// 对所有点进行dfs
	for vertex := range graph {
		// 已经访问过的节点不需要再进行dfs
		if visited[vertex] == 0 {
			traverse(graph, vertex, 1)
		}
	}
	return result
}

// 对节点进行dfs标记
func traverse(graph [][]int, vertex int, flag int) {
	// 判断节点是否已经被标记过
	if visited[vertex] != 0 {
		// 如果已经被标记过，并且和之前的标记结果不一致的话，则不是二分图
		if visited[vertex] != flag {
			result = false
		}
		return
	}
	// 相邻节点使用另一个状态进行标记
	visited[vertex] = flag
	if flag == 1 {
		flag = 2
	} else {
		flag = 1
	}
	// 遍历相邻节点进行标记
	for _, neighbor := range graph[vertex] {
		traverse(graph, neighbor, flag)
	}
}
