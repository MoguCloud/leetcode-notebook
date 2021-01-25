// Brute Force
//
// Your runtime beats 7.69 % of golang submissions (476 ms)
// Your memory usage beats 7.69 % of golang submissions (6.1 MB)
//
// 每个*号做左括号、右括号、空的尝试

func checkValidString(s string) bool {
	stack := []byte{}
	return dfs(0, &s, stack)
}

func dfs(i int, s *string, stack []byte) bool {
	if i == len(*s) {
		return len(stack) == 0
	}
	if (*s)[i] == '(' {
		stack = append(stack, '(')
		return dfs(i+1, s, stack)
	} else if (*s)[i] == ')' {
		if len(stack) == 0 {
			return false
		} else {
			stack = stack[:len(stack)-1]
			return dfs(i+1, s, stack)
		}
	} else {
		// * = '('
		stack1 := append(stack, '(')
		ret1 := dfs(i+1, s, stack1)
		// * = ''
		ret2 := dfs(i+1, s, stack)
		// * = ')'
		var ret3 bool
		if len(stack) == 0 {
			ret3 = false
		} else {
			stack2 := stack[:len(stack)-1]
			ret3 = dfs(i+1, s, stack2)
		}
		return ret1 || ret2 || ret3
	}
}
