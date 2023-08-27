// Your runtime beats 55.00 % of golang submissions (61 ms)
// Your memory usage beats 25.00 % of golang submissions (8.00 MB)
//
// 记忆化DFS
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n^2)

func canCross(stones []int) bool {
	visited := make(map[[2]int]struct{})
	var dfs func(from int, to int) bool
	dfs = func(from int, to int) bool {
		if to == len(stones)-1 {
			return true
		}
		key := [2]int{from, to}
		if _, ok := visited[key]; ok {
			return false
		}
		visited[key] = struct{}{}
		k := stones[to] - stones[from]
		ret := false
		for i := to + 1; i < len(stones) && stones[i] <= stones[to]+k+1; i++ {
			if stones[i] == stones[to]+k-1 {
				ret = dfs(to, i) || ret
			} else if stones[i] == stones[to]+k {
				ret = dfs(to, i) || ret
			} else if stones[i] == stones[to]+k+1 {
				ret = dfs(to, i) || ret
			}
		}
		return ret
	}
	if stones[1] != 1 {
		return false
	}
	return dfs(0, 1)
}
