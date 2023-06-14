// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 58.94 % of golang submissions (2 MB)
//
// 遍历
// 遍历字符串，针对每个字符，以该字符为中心向两边扩展，判断是否可以组成回文字符串
// 需要注意长度为奇数和长度为偶数的情况
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)

func countSubstrings(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		// 单个字母
		ret++
		// 向前扩展 长度为偶数
		for j, k := i-1, i; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] == s[k] {
				ret++
			} else {
				break
			}
		}
		// 长度为奇数
		for j, k := i-1, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] == s[k] {
				ret++
			} else {
				break
			}
		}
	}
	return ret
}
