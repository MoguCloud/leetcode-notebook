// Your runtime beats 91.67 % of golang submissions (115 ms)
// Your memory usage beats 91.67 % of golang submissions (7.5 MB)
//
// 数组 arr 用于存放 words2 中每个字符串中每个字符出现最大次数
// 再遍历 words1，words1 中的每个字符串的字符出现次数都大于等于 arr 中的次数，即为通用单词

func wordSubsets(words1 []string, words2 []string) []string {
	ret := []string{}
	arr := make([]int, 26)
	for _, word := range words2 {
		strMap := make(map[int]int)
		for _, str := range word {
			strMap[int(str)-'a'] += 1
		}
		for str, count := range strMap {
			if count > arr[str] {
				arr[str] = count
			}
		}
	}
	for _, word := range words1 {
		arrCopyed := make([]int, 26)
		copy(arrCopyed, arr)
		for _, str := range word {
			arrCopyed[int(str)-'a'] -= 1
		}
		if checkSum(arrCopyed) == 0 {
			ret = append(ret, word)
		}
	}
	return ret
}

func checkSum(arr []int) int {
	s := 0
	for _, n := range arr {
		if n > 0 {
			s += n
		}
	}
	return s
}
