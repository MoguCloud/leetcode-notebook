// Your runtime beats 88.1 % of golang submissions (30 ms)
// Your memory usage beats 23.81 % of golang submissions (6.5 MB)
//
// 遍历数组时，计算该时间取余数，并用数组记录该余数出现的次数，能与该时间凑成的满足条件的歌曲为余数+该余数=0或60。
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)


func numPairsDivisibleBy60(time []int) int {
	ret := 0
	timeCount := make([]int, 60)
	for _, t := range time {
		remainder := t % 60
		if remainder == 0 {
			ret += timeCount[0]
		} else {
			ret += timeCount[60-remainder]
		}
		timeCount[remainder] += 1
	}
	return ret
}
