// Your runtime beats 15.91 % of golang submissions (150 ms)
// Your memory usage beats 54.55 % of golang submissions (9.2 MB)
//
// hash
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func longestSubsequence(arr []int, difference int) int {
	memo := make(map[int]int) // 以key结尾子序列长度value
	ret := 1
	for _, num := range arr {
		length, _ := memo[num]
		length += 1
		memo[num+difference] = length
		if length > ret {
			ret = length
		}
	}
	return ret
}
