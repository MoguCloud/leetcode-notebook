// 贪心
//
// Your runtime beats 100.00 % of golang submissions (0 ms)
// Your memory usage beats 23.08 % of golang submissions (1.9 MB)
//
// minLeft, maxLeft 两个变量分别记录左括号可能出现的最小次数和最大次数。
// 当出现 ( 时，minLeft 和 maxLeft 都 +1
// 当出现 ) 时，minLeft 和 maxLeft 都 -1。如果 maxLeft <= 0，则不能匹配
// 当出现 * 时，minLeft -1, maxLeft +1

func checkValidString(s string) bool {
	maxLeft := 0
	minLeft := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			maxLeft++
			minLeft++
		} else if s[i] == ')' {
			if minLeft > 0 {
				minLeft--
			}
			if maxLeft > 0 {
				maxLeft--
			} else {
				return false
			}
		} else {
			if minLeft > 0 {
				minLeft--
			}
			maxLeft++
		}
	}
	if minLeft == 0 {
		return true
	} else {
		return false
	}
}
