// Your runtime beats 57.14 % of golang submissions (144 ms)
// Your memory usage beats 71.43 % of golang submissions (9.5 MB)
//
// Stack
// 将 path 按 '/' 分割，分出来的会是目录名、空字符串、一个点和两个点
// 遍历分割后的字符串数组，如果是目录名称就压入栈中；如果是空字符串就什么都不做；如果是一个点也什么都不做；如果是两个点就将栈顶元素弹出
// 然后将栈的元素用 '/' 连接，注意此时还少了最前面一个 '/' 需要补上
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func simplifyPath(path string) string {
	paths := strings.Split(path[1:], "/")
	stack := []string{}
	for _, p := range paths {
		if p == "." || p == "" {

		} else if p == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, p)
		}
	}
	return "/" + strings.Join(stack, "/")
}
