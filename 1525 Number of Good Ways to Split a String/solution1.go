// Your runtime beats 67.92 % of golang submissions (24 ms)
// Your memory usage beats 84.91 % of golang submissions (3.6 MB)
//
// Hash
//
// 2个哈希表，初始状态hash1为空，hash2为输入s中每个字母的数量
// 分割线i从s的左边到右边开始遍历，每向右边移动1位，hash1[s[i]]++, hash2[s[i]]--
// 如果 hash1 key的个数 和 hash2 key的个数相同时，则为好分割
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func numSplits(s string) int {
	ret := 0
	numCount := make(map[rune]int)
	for _, str := range s {
		numCount[str] += 1
	}
	prevNumCount := make(map[rune]int)
	for _, str := range s {
		prevNumCount[str] += 1
		numCount[str] -= 1
		if count, _ := numCount[str]; count == 0 {
			delete(numCount, str)
		}
		if len(numCount) == len(prevNumCount) {
			ret += 1
		}
	}
	return ret
}
