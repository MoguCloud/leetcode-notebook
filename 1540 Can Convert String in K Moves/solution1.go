// Your runtime beats 100 % of golang submissions (22 ms)
// Your memory usage beats 50 % of golang submissions (7.1 MB)
//
// 统计操作次数
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}
	// 用于存储变换1-25次的个数
	diff := make([]int, 26)
	for i := 0; i < len(s); i++ {
		// 当字符不同时，记录需要操作的次数
		if s[i] != t[i] {
			diff[(t[i]+26-s[i])%26]++
		}
	}
	move := 0
	for i := 0; i < 26; i++ {
		if diff[i] > 0 {
			// 当需要变换diff[i]次的个数>1时，每多一个需要增加26
			move = max(i+(diff[i]-1)*26, i)
			if move > k {
				return false
			}
		}
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
