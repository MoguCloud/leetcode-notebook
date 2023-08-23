// Your runtime beats 72.55 % of golang submissions (1 ms)
// Your memory usage beats 100 % of golang submissions (2 MB)
//
// 计数
// 当一个字母出现的次数大于 (len(s) + 1 ) / 2 时，则不满足条件
// 根据字母出现的次数进行排序
// 因为奇数位(下标0, 2, ...)的数量大于等于偶数位(下标1,3,5...)
// 所以将出现次数最多的字母全部插入奇数位即可保证不会出现相邻的情况
// 所以只要根据字母的数量依次插入奇数位，再依次插入偶数位即可
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func reorganizeString(s string) string {
	// 统计各字母出现的次数
	charCount := make([][2]int, 26)
	for i := 0; i < 26; i++ {
		charCount[i] = [2]int{i, 0}
	}
	for i := 0; i < len(s); i++ {
		charCount[s[i]-'a'][1]++
	}
	// 判断出现最多的字母是否大于最大值
	maxCount := 0
	for i := 0; i < 26; i++ {
		count := charCount[i][1]
		if count > maxCount {
			maxCount = count
		}
	}
	// 不满足条件直接返回
	if maxCount > (len(s)+1)/2 {
		return ""
	}
	// 按字母个数降序排序
	sort.Slice(charCount, func(i, j int) bool { return charCount[i][1] > charCount[j][1] })
	// 初始化返回值
	ret := make([]byte, len(s))
	// j 用于记录需要插入的字母的索引
	j := 0
	// 根据字母个数插入到奇数位
	for i := 0; i < len(ret); i = i + 2 {
		// 判断当前字母是否全部插入完毕，全部插入完毕则处理下一个字母
		for charCount[j][1] <= 0 {
			j++
		}
		ret[i] = byte(charCount[j][0] + 'a')
		charCount[j][1]--
	}
	// 根据字母个数插入到偶数位
	for i := 1; i < len(ret); i = i + 2 {
		// 判断当前字母是否全部插入完毕，全部插入完毕则处理下一个字母
		for charCount[j][1] <= 0 {
			j++
		}
		ret[i] = byte(charCount[j][0] + 'a')
		charCount[j][1]--
	}
	return string(ret)
}
