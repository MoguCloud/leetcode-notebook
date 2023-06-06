// Your runtime beats 94.12 % of golang submissions (277 ms)
// Your memory usage beats 76.47 % of golang submissions (48.3 MB)
//
// BFS
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(m)

func minimumOperations(nums []int, start int, goal int) int {
	memo := make(map[int]struct{})
	q := []int{start}
	count := 0
	for len(q) > 0 {
		newQ := []int{}
		for _, n := range q {
			if n == goal {
				return count
			}
			if n < 0 || n > 1000 {
				continue
			}
			if _, ok := memo[n]; ok {
				continue
			}
			memo[n] = struct{}{}
			for _, num := range nums {
				newQ = append(newQ, n+num)
				newQ = append(newQ, n-num)
				newQ = append(newQ, n^num)
			}
		}
		count++
		q = newQ
	}
	return -1
}
