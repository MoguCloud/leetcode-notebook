// Your runtime beats 76.47 % of golang submissions (4 ms)
// Your memory usage beats 11.76 % of golang submissions (3.8 MB)
//
// Stack
// 模拟栈，一步一步模拟将pushed数组压入栈中，如果栈顶元素等于popped数组的话，则弹出
// 如果遍历完成后，栈为空则可以满足条件
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _, p := range pushed {
		stack = append(stack, p)
		for len(stack) > 0 {
			if stack[len(stack)-1] == popped[i] {
				stack = stack[:len(stack)-1]
				i++
			} else {
				break
			}
		}
	}
	return len(stack) == 0
}
