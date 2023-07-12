// Your runtime beats 88.73 % of golang submissions (99 ms)
// Your memory usage beats 61.97 % of golang submissions (7.4 MB)
//
// 反向图+拓扑排序
// 如果一个点的出度为0，则该节点是安全的。如果一个节点可以到达安全节点的话，那该节点也是安全的
// 即相应的反向图进行拓扑排序，排序完成后反向图的入度（原图的出度）为0的点就是安全的点
//
// 时间复杂度 O(E+V)
// 空间复杂度 O(E+V)

func eventualSafeNodes(graph [][]int) []int {
	reverseGraph := make([][]int, len(graph)) // 反向图
	outdegree := make([]int, len(graph))      // 原图的出度，反向图的入度
	// 建立反向图
	for i := 0; i < len(reverseGraph); i++ {
		reverseGraph[i] = []int{}
	}
	for i, nodes := range graph {
		outdegree[i] = len(nodes)
		for _, node := range nodes {
			reverseGraph[node] = append(reverseGraph[node], i)
		}
	}
	// 根据反向图拓扑排序
	q := []int{}
	for i, d := range outdegree {
		if d == 0 {
			q = append(q, i)
		}
	}
	for i := 0; i < len(q); i++ {
		nodes := reverseGraph[q[i]]
		for _, node := range nodes {
			outdegree[node]--
			if outdegree[node] == 0 {
				q = append(q, node)
			}
		}
	}
	// 返回出度为0的节点
	ret := []int{}
	for i, d := range outdegree {
		if d == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}
