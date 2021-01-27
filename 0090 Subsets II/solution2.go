// Backtrack
//
// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (2.4 MB)
//
// 先排序，回溯的时候判断是否与上一个数是否相同。
// 以 [1, 2, 2, 3]举例
// 回溯过程如下
// ![](https://imgur.com/w4dKTFi)   
// 即当前节点与长兄节点相同时，则剪枝

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	tmp := []int{}
	backtrack(&nums, 0, &ret, &tmp)
	return ret
}

func backtrack(nums *[]int, idx int, ret *[][]int, tmp *[]int) {
	copyedTmp := make([]int, len(*tmp))
	copy(copyedTmp, *tmp)
	*ret = append(*ret, copyedTmp)
	if idx >= len(*nums) {
		return
	}
	prev := 0
	for start := idx; start < len(*nums); start++ {
		if start != idx && (*nums)[start] == prev {
			continue
		}
		*tmp = append(*tmp, (*nums)[start])
		backtrack(nums, start+1, ret, tmp)
		*tmp = (*tmp)[:len(*tmp)-1]
		prev = (*nums)[start]
	}
}
