// Your runtime beats 88.61 % of golang submissions (8 ms)
// Your memory usage beats 41.77 % of golang submissions (7.8 MB)
//
// 回溯

var ret [][]int

func combine(n int, k int) [][]int {
	ret = [][]int{}
	backtrack(n, 1, k, []int{})
	return ret
}

func backtrack(n int, start int, k int, tmp []int) {
	// 终止条件是遍历到长度为k
	if k == 0 {
		ans := make([]int, len(tmp))
		copy(ans, tmp)
		ret = append(ret, ans)
		return
	}
	// 从 start 作为起点开始搜索，避免重复
	for i := start; i <= n; i++ {
		tmp = append(tmp, i)
		// 下一轮搜索起点+1，避免出现重复情况
		backtrack(n, i+1, k-1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}
