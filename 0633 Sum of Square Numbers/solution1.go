// Your runtime beats 11.11 % of golang submissions
// Your memory usage beats 11.11 % of golang submissions (7.1 MB)
//
// 创建 [0, sqrt(c) ]的哈希表sqr
// 遍历哈希表看哈希表中是否存在 c - sqrt(c)
//
// 时间复杂度 O(sqrt(c))
// 空间复杂度 O(sqrt(c))
func judgeSquareSum(c int) bool {
	ret := false
	sqr := make(map[int]struct{})
	i := 0
	for {
		s := i * i
		if s > c {
			break
		}
		sqr[s] = struct{}{}
		i += 1
	}
	for i, _ := range sqr {
		if _, ok := sqr[c-i]; ok {
			ret = true
			break
		}
	}
	return ret
}
