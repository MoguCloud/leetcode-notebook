// Your runtime beats 71 % of golang submissions (8 ms)
// Your memory usage beats 74.51 % of golang submissions (2.7 MB)
//
// 滑动窗口
// 用两个指针分别指向子串的左右边界。
// 右指针不断向右移动，当右指针指向的元素出现在子串中（即子串中有重复元素），则将左指针不断向又移动，直到指向重复元素的下一个元素（即子串中没有重复元素）。
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func lengthOfLongestSubstring(s string) int {
	strMap := make(map[byte]int)
	i := 0
	ret := 0
	subStrLen := 0
	for j := 0; j < len(s); j++ {
		idx, ok := strMap[s[j]]
		if ok {
			subStrLen = j - i
			for k := i; k <= idx; k++ {
				delete(strMap, s[k])
			}
			i = idx + 1
			strMap[s[j]] = j
		} else {
			subStrLen = j - i + 1
			strMap[s[j]] = j
		}
		if subStrLen > ret {
			ret = subStrLen
		}
	}
	return ret
}
