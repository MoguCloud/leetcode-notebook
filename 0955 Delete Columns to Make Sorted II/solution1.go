// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 20.00 % of golang submissions (2.8 MB)
//
// 对于第j列来说
// strs[i-1][j] > strs[i][j] 时，是必须删除的，
// strs[i-1][j] < strs[i][j] 时，是不需要删除的，并且对于任意列来说strs[i-1]都小于strs[i]
// strs[i-1][j] = strs[i][j] 时，需要判断前j-1列，strs[i-1]是否小于strs[i]。
//
// 使用一个数组visited，visited[i]用来记录strs[i-1]是否小于strs[i]
// 需要注意的是，在遍历的时候，需要在遍历完该列后再判断strs[i-1]是否小于strs[i]，
// 因为如果删除了该列的话，再以该列的大小来判断是错误的。
//
// 时间复杂度 O(mn)
// 空间复杂度 O(n)

func minDeletionSize(strs []string) int {
	ret := 0
	n := len(strs)
	m := len(strs[0])
	visited := make([]bool, n)
	for i := 0; i < m; i++ {
		isDeleted := false
		for j := 1; j < n; j++ {
			if visited[j] {
				continue
			}
			if strs[j][i] < strs[j-1][i] {
				ret += 1
				isDeleted = true
				break
			}
		}
		if !isDeleted {
			for j := 1; j < n; j++ {
				if strs[j][i] > strs[j-1][i] {
					visited[j] = true
				}
			}
		}
	}
	return ret
}
