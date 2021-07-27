// Your runtime beats 94.87 % of golang submissions (4 ms)
// Your memory usage beats 74.36 % of golang submissions (3.2 MB)
//
// 因为 0 <= trips[i][1] < trips[i][2] <= 1000，站数最多为1000
// 所以可以创建一个长度为1000的数组，记录每一个站的上下车人数
// 对这个数组累计求和，只要出现 > capacity时，则该站的乘客数会超过容量，返回false
// 若遍历完都没有 > capacity，则返回true
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func carPooling(trips [][]int, capacity int) bool {
	pick := make([]int, 1000)
	for i := 0; i < len(trips); i++ {
		pick[trips[i][1]] += trips[i][0]
		pick[trips[i][2]] -= trips[i][0]
	}
	passengers := 0
	for i := 0; i < len(pick); i++ {
		passengers += pick[i]
		if passengers > capacity {
			return false
		}
	}
	return true
}
