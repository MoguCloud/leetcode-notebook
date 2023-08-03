// Your runtime beats 70.83 % of golang submissions (1 ms)
// Your memory usage beats 51.17 % of golang submissions (2.06 MB)
//
// 回溯
//
// 时间复杂度 O(3^m * 4^n)
// 空间复杂度 O(m+n)
// m为2-6和8的数量 n为7和9的数量

var m []string = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	ret := []string{}
	if len(digits) == 0 {
		return ret
	}
	var backtrack func(ans []byte)
	backtrack = func(ans []byte) {
		if len(ans) == len(digits) {
			ret = append(ret, string(ans))
			return
		}
		curr := int(digits[len(ans)] - '2')
		for i := 0; i < len(m[curr]); i++ {
			ans = append(ans, m[curr][i])
			backtrack(ans)
			ans = ans[:len(ans)-1]
		}
	}
	backtrack([]byte{})
	return ret
}
