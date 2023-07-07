// Your runtime beats 55.56 % of golang submissions (16 ms)
// Your memory usage beats 88.89 % of golang submissions (5.3 MB)
//
// 滑动窗口
// 这题可以理解成包含另一个字符个数不超过k的最长连续个数
// 即包含不超过k个F的连续T长度 和 包含不超过k个T的连续F长度 取最大值
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func maxConsecutiveAnswers(answerKey string, k int) int {
	maxChr := func(chr byte, k int) int {
		left := 0
		right := 0
		ret := 0
		for right < len(answerKey) {
			if answerKey[right] == chr {
				if k == 0 {
					for left < right {
						if answerKey[left] == chr {
							k++
							left++
							break
						}
						left++
					}
				} else {
					k--
					right++
				}
			} else {
				right++
			}
			if right-left > ret {
				ret = right - left
			}
		}
		return ret
	}
	return max(maxChr('T', k), maxChr('F', k))
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
