// Your runtime beats 39.61 % of golang submissions (13 ms)
// Your memory usage beats 93.72 % of golang submissions (4.6 MB)
//
// 栈模拟
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for i := 0; i < len(asteroids); i++ {
		if len(stack) == 0 {
			// 栈为空时 入栈
			stack = append(stack, asteroids[i])
		} else if (stack[len(stack)-1] > 0 && asteroids[i] > 0) || (stack[len(stack)-1] < 0 && asteroids[i] < 0) || (stack[len(stack)-1] < 0 && asteroids[i] > 0) {
			// 当前行星与上一颗行星同向或者背向时 入栈
			stack = append(stack, asteroids[i])
		} else {
			// 当前行星与上一颗存活的行星相向时 进行大小判断是否存活
			finished := false
			for len(stack) > 0 {
				if (stack[len(stack)-1] > 0 && asteroids[i] > 0) || (stack[len(stack)-1] < 0 && asteroids[i] < 0) || (stack[len(stack)-1] < 0 && asteroids[i] > 0) {
					break
				}
				if abs(stack[len(stack)-1]) == abs(asteroids[i]) {
					stack = stack[:len(stack)-1]
					finished = true
					break
				}
				if abs(stack[len(stack)-1]) > abs(asteroids[i]) {
					finished = true
					break
				}
				if abs(stack[len(stack)-1]) < abs(asteroids[i]) {
					stack = stack[:len(stack)-1]
					continue
				}
			}
			if !finished {
				stack = append(stack, asteroids[i])
			}
		}
	}
	return stack
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
