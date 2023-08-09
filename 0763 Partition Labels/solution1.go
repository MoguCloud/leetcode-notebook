// Your runtime beats 75.26 % of golang submissions (2 ms)
// Your memory usage beats 98.97 % of golang submissions (2.09 MB)
//
// 贪心
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func partitionLabels(s string) []int {
	ret := []int{}
	// 记录每个字母最后出现的位置
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a'] = i
	}
	i := 0
	j := 0
	k := 0
	// 遍历字符串s
	for k < len(s) {
		// i 为划分字符串的左边界
		i = k
		// j 为划分字符串的右边界
		j = m[s[i]-'a']
		// 遍历s[i+1:j]
		for k = i + 1; k < j; k++ {
			// 因为出现过的字符串必须在一个片段中
			// 所以右边界扩大到出现的字母的最后出现的位置
			if m[s[k]-'a'] > j {
				j = m[s[k]-'a']
			}
		}
		k = j + 1
		// 划分字符串的长度为 j - i + 1
		ret = append(ret, j-i+1)
	}
	return ret
}
