// HashMap
//
// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 91.67 % of golang submissions (2.8 MB)
//
// 通过哈希表存每个数 (num) 出现的次数 (numCount) 。
// num 可以按 (num + 1)个构成一个分组。 「我说有 2 个人颜色和我一样是红色，即红色有 2+1 个人」
// numCount 可以拆分成 numCount / (num + 1) 向上取整个分组。
// 特殊情况 num = 0 时，单独构成一个分组，该分组里有 1 人。

func numRabbits(answers []int) int {
	countMap := make(map[int]int)
	ret := 0
	for i := 0; i < len(answers); i++ {
		if answers[i] == 0 {
			ret += 1
		} else {
			countMap[answers[i]] += 1
		}
	}
	for k, v := range countMap {
		ret += int(math.Ceil(float64(v)/float64(k+1))) * (k + 1)
	}
	return ret
}
