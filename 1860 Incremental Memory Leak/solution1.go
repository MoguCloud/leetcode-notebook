// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (2 MB)
//
// 模拟法

func memLeak(memory1 int, memory2 int) []int {
	t := 1
	for t <= max(memory1, memory2) {
		if memory1 >= memory2 {
			memory1 -= t
		} else {
			memory2 -= t
		}
		t += 1
	}
	return []int{t, memory1, memory2}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
