// Your runtime beats 93.59 % of golang submissions (7 ms)
// Your memory usage beats 96.15 % of golang submissions (4.8 MB)
//
// 滑动窗口
// 先计算初始窗口[0:k]中的元音个数
// 然后将窗口将右侧滑动
// 当左侧移出窗口的元素是元音时，元音个数-1
// 当右侧进入窗口的元素是元音时，元音个数+1
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func maxVowels(s string, k int) int {
	n := 0
	i := 0
	for i = 0; i < k; i++ {
		if IsVowel(s[i]) {
			n++
		}
	}
	ret := n
	for i < len(s) {
		if IsVowel(s[i-k]) {
			n--
		}
		if IsVowel(s[i]) {
			n++
		}
		if n > ret {
			ret = n
		}
		i++
	}
	return ret
}

func IsVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
