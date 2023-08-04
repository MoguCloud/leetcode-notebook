// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 44.06 % of golang submissions (2.25 MB)
//
// 记忆化回溯
//
// 时间复杂度 O(len(s) * len(wordDict))
// 空间复杂度 O(len(s))

func wordBreak(s string, wordDict []string) bool {
	memo := make(map[int]bool)
	var backtrack func(offset int) bool
	backtrack = func(offset int) bool {
		// 回溯递归终止条件
		if offset >= len(s) {
			return true
		}
		// 判断是否计算过，避免重复计算
		if ret, ok := memo[offset]; ok {
			return ret
		}
		for i := 0; i < len(wordDict); i++ {
			// 判断 wordDict 中的字符串是否和s匹配
			if s[offset] == wordDict[i][0] && offset+len(wordDict[i]) <= len(s) && wordDict[i] == s[offset:offset+len(wordDict[i])] {
				// 当后续全部满足时才返回，如果当前状态返回false则需要继续遍历
				if backtrack(offset + len(wordDict[i])) {
					return true
				}
			}
		}
		// 记录计算结果
		memo[offset] = false
		return false
	}
	return backtrack(0)
}
