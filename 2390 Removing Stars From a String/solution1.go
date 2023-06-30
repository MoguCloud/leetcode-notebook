// Your runtime beats 59.26 % of golang submissions (29 ms)
// Your memory usage beats 55.56 % of golang submissions (7.3 MB)
//
// Stack
// 遇到非星号的字符往栈里压，遇到星号弹出栈顶元素
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func removeStars(s string) string {
	stack := []rune{}
	for _, str := range s {
		if str != '*' {
			stack = append(stack, str)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
