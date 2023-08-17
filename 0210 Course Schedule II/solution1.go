// Your runtime beats 99.70 % of golang submissions (2 ms)
// Your memory usage beats 96.04 % of golang submissions (6.17 MB)
//
// 拓扑排序
//
// 时间复杂度 O(n)
// 空间复杂度 O(m*n)

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 入度
	indegrees := make([]int, numCourses)
	// 邻接表
	graph := make([][]int, numCourses)
	// 计算入度和邻接表
	for _, req := range prerequisites {
		indegrees[req[0]]++
		graph[req[1]] = append(graph[req[1]], req[0])
	}
	order := []int{}
	for course, indegree := range indegrees {
		if indegree == 0 {
			order = append(order, course)
		}
	}
	// 拓扑排序
	for i := 0; i < len(order); i++ {
		course := order[i]
		for _, nextCourse := range graph[course] {
			indegrees[nextCourse]--
			if indegrees[nextCourse] == 0 {
				order = append(order, nextCourse)
			}
		}
	}
	// 当能够完成所有课程时才返回结果数组
	if len(order) == numCourses {
		return order
	}
	// 否则返回空数组
	return []int{}
}
