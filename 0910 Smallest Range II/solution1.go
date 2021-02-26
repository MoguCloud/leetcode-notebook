// 贪心
//
// Your runtime beats 80 % of golang submissions (24 ms)
// Your memory usage beats 93.14 % of golang submissions (5.9 MB)
//
// https://leetcode-cn.com/problems/smallest-range-ii/solution/tai-nan-liao-zhi-neng-hua-tu-ping-zhi-jue-by-user8/

func smallestRangeII(A []int, K int) int {
	sort.Ints(A)
	minD := A[len(A)-1] - A[0]
	for i := 0; i < len(A)-1; i++ {
		ma := max(A[i]+K, A[len(A)-1]-K)
		mi := min(A[0]+K, A[i+1]-K)
		d := ma - mi
		minD = min(minD, d)
	}
	return minD
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
