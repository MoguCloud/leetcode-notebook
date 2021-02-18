// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 27.27 % of golang submissions (2 MB)
//
// 如果N次指令完，面朝北，位置再原点，则成环，如果不在原点就不存在环。
// 一次指令完成后可能出现的方向
// 1. 没有改变
// 2. 顺时针90度
// 3. 逆时针90度
// 4. 180度
// 这4种情况再执行4次后都能回到最一开始的方向，也就是朝北。
// 所以运行4次指令后，在原点则成环，否则就不成环。
//
// L (1, 0) -> (0, -1) -> (-1, 0) -> (0, 1)
// R (1, 0) -> (0, 1) -> (-1, 0) -> (0, -1)

func isRobotBounded(instructions string) bool {
	direction := []int{1, 0}
	pos := []int{0, 0}
	for i := 0; i < 4; i++ {
		for j := 0; j < len(instructions); j++ {
			if instructions[j] == 'G' {
				pos[0] += direction[0]
				pos[1] += direction[1]
			} else {
				if direction[0] == 1 && direction[1] == 0 {
					if instructions[j] == 'L' {
						direction = []int{0, -1}
					} else {
						direction = []int{0, 1}
					}
				} else if direction[0] == 0 && direction[1] == -1 {
					if instructions[j] == 'L' {
						direction = []int{-1, 0}
					} else {
						direction = []int{1, 0}
					}
				} else if direction[0] == -1 && direction[1] == 0 {
					if instructions[j] == 'L' {
						direction = []int{0, 1}
					} else {
						direction = []int{0, -1}
					}
				} else if direction[0] == 0 && direction[1] == 1 {
					if instructions[j] == 'L' {
						direction = []int{1, 0}
					} else {
						direction = []int{-1, 0}
					}
				}
			}
		}
	}
	if pos[0] == 0 && pos[1] == 0 {
		return true
	}
	return false
}
