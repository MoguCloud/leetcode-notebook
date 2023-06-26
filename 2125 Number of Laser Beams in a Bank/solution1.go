// Your runtime beats 88.46 % of golang submissions (25 ms)
// Your memory usage beats 80.77 % of golang submissions (7.3 MB)
//
// 2行中间的激光数量 = 上一次的安全设备数量 * 下一层的安全设备数量
//
// 时间复杂度 O(m*n)
// 空间复杂度 O(1)

func numberOfBeams(bank []string) int {
	ret := 0
	i := 0
	prev := 0
	curr := 0
	for i < len(bank) {
		curr = 0
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				curr++
			}
		}
		if curr > 0 {
			ret += curr * prev
			prev = curr
		}
		i++
	}
	return ret
}
