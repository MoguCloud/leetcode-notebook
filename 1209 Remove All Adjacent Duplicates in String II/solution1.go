// Your runtime beats 96.15 % of golang submissions (4 ms)
// Your memory usage beats 69.23 % of golang submissions (6.6 MB)
//
// 栈 strStack 存放倒序的 s
// 栈 countStack 存放倒序的 s 不同字母出现的次序
// 倒序遍历 s，当countStack栈顶元素等于k时，
// 也就是出现连续k个相同字母时候，
// 从strStack弹出k个元素(消除k个相同元素)
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func removeDuplicates(s string, k int) string {
	strStack := []byte{}
	countStack := []int{}
	var prevByte byte
	for i := len(s) - 1; i >= 0; i-- {
		strStack = append(strStack, s[i])
		if prevByte == s[i] {
			countStack[len(countStack)-1] += 1
			if countStack[len(countStack)-1] == k {
				handleRemove(&strStack, &countStack, k)
				if len(strStack) > 0 {
					prevByte = strStack[len(strStack)-1]
				} else {
					prevByte = 0
				}
			}
		} else {
			prevByte = s[i]
			countStack = append(countStack, 1)
		}
	}
	for i, j := 0, len(strStack)-1; i < j; i, j = i+1, j-1 {
		strStack[i], strStack[j] = strStack[j], strStack[i]
	}
	return string(strStack)
}

func handleRemove(strStack *[]byte, countStack *[]int, k int) {
	if (*countStack)[len(*countStack)-1] != k {
		return
	}
	*strStack = (*strStack)[:len(*strStack)-k]
	*countStack = (*countStack)[:len(*countStack)-1]
}
