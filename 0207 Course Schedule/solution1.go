// Your runtime beats 66.99 % of golang submissions (12 ms)
// Your memory usage beats 75.73 % of golang submissions (6.3 MB)
//
// 拓扑排序
//
// 时间复杂度 O(E+V)
// 空间复杂度 O(E+V)

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化图和入度
	indegree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, req := range prerequisites {
		graph[req[1]] = append(graph[req[1]], req[0])
		indegree[req[0]]++
	}
	q := []int{}
	for i, d := range indegree {
		if d == 0 {
			q = append(q, i)
		}
	}
	// 拓扑排序
	for i := 0; i < len(q); i++ {
		vertexs := graph[q[i]]
		for _, v := range vertexs {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(q) == numCourses
}
