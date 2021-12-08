// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 32 % of golang submissions (2.5 MB)
//
// 模拟法
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func findTheWinner(n int, k int) int {
	circle := []int{}
	for i := 0; i < n; i++ {
		circle = append(circle, i+1)
	}
	offset := 0
	for len(circle) > 1 {
		offset = (offset - 1 + k) % len(circle)
		circle = append(circle[0:offset], circle[offset+1:len(circle)]...)
	}
	return circle[0]
}
