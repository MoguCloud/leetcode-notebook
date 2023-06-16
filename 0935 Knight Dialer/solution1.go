// Your runtime beats 7.14 % of golang submissions (578 ms)
// Your memory usage beats 7.14 % of golang submissions (13.1 MB)
//
// 记忆化搜索
//
// 时间复杂度 O(10^n)
// 空间复杂度 O(n)

var next map[int][]int = map[int][]int{
	1: []int{8, 6},
	2: []int{7, 9},
	3: []int{4, 8},
	4: []int{3, 9, 0},
	5: []int{},
	6: []int{7, 1, 0},
	7: []int{2, 6},
	8: []int{1, 3},
	9: []int{2, 4},
	0: []int{4, 6},
}
var memo map[[2]int]int

func knightDialer(n int) int {
	ret := 0
	memo = make(map[[2]int]int)
	for i := 0; i <= 9; i++ {
		res := dfs(i, 1, n)
		res %= 1e9 + 7
		ret += res
		ret %= 1e9 + 7
	}
	return ret
}

func dfs(cur int, count int, total int) int {
	if count == total {
		return 1
	}
	key := [2]int{cur, count}
	ret, ok := memo[key]
	if ok {
		return ret
	}
	ret = 0
	for _, n := range next[cur] {
		res := dfs(n, count+1, total)
		res %= (1e9 + 7)
		ret += res
		ret %= 1e9 + 7
	}
	memo[key] = ret
	return ret
}
