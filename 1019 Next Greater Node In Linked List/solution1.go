// Your runtime beats 84 % of golang submissions (56 ms)
// Your memory usage beats 36 % of golang submissions (7.7 MB)
//
// 单调栈
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func nextLargerNodes(head *ListNode) []int {
	ret := []int{}
	for node := head; node != nil; node = node.Next {
		ret = append(ret, node.Val)
	}
	s := []int{}
	for i := len(ret) - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= ret[i] {
			s = s[:len(s)-1]
		}
		tmp := ret[i]
		if len(s) == 0 {
			ret[i] = 0
		} else {
			ret[i] = s[len(s)-1]
		}
		s = append(s, tmp)
	}
	return ret
}
