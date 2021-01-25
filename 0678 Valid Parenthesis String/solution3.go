// 双栈
// 
// Your runtime beats 100.00 % of golang submissions (0 ms)
// Your memory usage beats 23.08 % of golang submissions (1.9 MB)
//
// 两个栈 leftStack starStack 分别存放 左括号和星号的索引
// 当遍历到右括号时，leftStack出栈
// 遍历完后，看 starStack 的长度是否大于 leftStack。
// 如果 starStack 小于 leftStack，则有多的左括号，返回 false
// starStack 和 leftStack 各弹出一个，看 star 的索引是否大于 left 的，如果 star 索引小于 left 索引，则该 left 没有能够匹配的星号/右括号，返回 false


func checkValidString(s string) bool {
	leftStack := []int{}
	starStack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftStack = append(leftStack, i)
		} else if s[i] == ')' {
			if len(leftStack) > 0 {
				leftStack = leftStack[:len(leftStack)-1]
			} else {
				if len(starStack) > 0 {
					starStack = starStack[:len(starStack)-1]
				} else {
					return false
				}
			}
		} else {
			starStack = append(starStack, i)
		}
	}
	if len(leftStack) > len(starStack) {
		return false
	}
	for len(leftStack) > 0 {
		leftIdx := leftStack[len(leftStack)-1]
		starIdx := starStack[len(starStack)-1]
		if leftIdx > starIdx {
			return false
		}
		leftStack = leftStack[:len(leftStack)-1]
		starStack = starStack[:len(starStack)-1]
	}
	return true
}
