// 哈希表 + 滑动窗口
//
// Your runtime beats 92.86 % of golang submissions (64 ms)
// Your memory usage beats 50 % of golang submissions (7.1 MB)
//
// 时间复杂度 O(n)
// 空间复杂度 O(2 ** k)
//
// 将字符串 s 的所有长度为 k 的字串存入哈希表中，如果存入的个数为 2 ** k 次，则满足条件可以提前退出。
// 因为 s 每一位都是 0 和 1 ，所以可以将哈希表的 key 设置成 int 类型，进一步节省空间。

func hasAllCodes(s string, k int) bool {
	set := make(map[int]struct{})
	count := 0
	for i := 0; i <= len(s)-k; i++ {
		n, _ := strconv.Atoi(s[i : i+k])
		_, ok := set[n]
		if !ok {
			set[n] = struct{}{}
			count += 1
		}
		if count == (1 << k) {
			break
		}
	}
	return count == (1 << k)
}
