// Your runtime beats 79.31 % of golang submissions (8 ms)
// Your memory usage beats 92.24 % of golang submissions (6.5 MB)
//
// DFS
//
// 时间复杂度 O(n*2^n)
// 空间复杂度 O(n)

func allPathsSourceTarget(graph [][]int) [][]int {
	ret := [][]int{}
	var dfs func(from int, order []int)
	dfs = func(from int, order []int) {
		order = append(order, from)
		if from == len(graph)-1 {
			ans := make([]int, len(order))
			copy(ans, order)
			ret = append(ret, ans)
		}
		for _, dst := range graph[from] {
			dfs(dst, order)
		}
		order = order[:len(order)-1]
	}
	dfs(0, []int{})
	return ret
}
