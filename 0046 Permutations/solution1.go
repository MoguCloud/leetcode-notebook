// Your runtime beats 61.28 % of golang submissions (2 ms)
// Your memory usage beats 70.14 % of golang submissions (2.68 MB)
//
// 回溯
//
// 时间复杂度 O(n*n!)
// 空间复杂度 O(n)

func permute(nums []int) [][]int {
	ret := [][]int{}
	var backtrack func(flag int, ans []int)
	backtrack = func(flag int, ans []int) {
		if len(ans) == len(nums) {
			tmp := make([]int, len(ans))
			copy(tmp, ans)
			ret = append(ret, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			// flag 用于判断某个数是否被使用过，避免重复使用
			if flag&(1<<i) == 0 {
				ans = append(ans, nums[i])
				backtrack(flag|(1<<i), ans)
				ans = ans[:len(ans)-1]
			}
		}
	}
	backtrack(0, []int{})
	return ret
}
