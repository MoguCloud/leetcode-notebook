// Your runtime beats 67.4 % of golang submissions (4 ms)
// Your memory usage beats 31.92 % of golang submissions (2.6 MB)
//
// 滑动窗口
// 统计s1中各个字母出现的次数
// 取一个长度为len(s1)的窗口不断遍历s2，统计窗口中各个字母出现的次数，
// 如果和s1的相同，则返回true，如果遍历完也没有，则返回false
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	var arr1 [26]int
	var arr2 [26]int
	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}
	if arr1 == arr2 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		arr2[s2[i]-'a']++
		arr2[s2[i-len(s1)]-'a']--
		if arr1 == arr2 {
			return true
		}
	}
	return false
}
