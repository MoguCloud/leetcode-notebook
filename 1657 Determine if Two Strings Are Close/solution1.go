// Your runtime beats 55 % of golang submissions (80 ms)
// Your memory usage beats 5 % of golang submissions (7.7 MB)
//
// 操作1 可以推出 word1 和 word2 中的字母种类要一样
// 操作2 可以推出 word1 和 word2 中字母分组数要一样(如word1中要有1个a 2个b， word2中要有2个c 1个d)
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	countMap1 := make(map[rune]int)
	countMap2 := make(map[rune]int)
	for _, str := range word1 {
		countMap1[str] += 1
	}
	for _, str := range word2 {
		countMap2[str] += 1
	}
	countArr1 := []int{}
	countArr2 := []int{}
	for _, v := range countMap1 {
		countArr1 = append(countArr1, v)
	}
	for k, v := range countMap2 {
		countArr2 = append(countArr2, v)
		if _, ok := countMap1[k]; !ok {
			return false
		}
	}
	sort.Ints(countArr1)
	sort.Ints(countArr2)
	if len(countArr1) != len(countArr2) {
		return false
	}
	for i := 0; i < len(countArr1); i++ {
		if countArr1[i] != countArr2[i] {
			return false
		}
	}
	return true
}
