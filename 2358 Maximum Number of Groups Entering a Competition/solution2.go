// Your runtime beats 50 % of golang submissions (88 ms)
// Your memory usage beats 50 % of golang submissions (8.9 MB)
//
// 1个最小的数 必然小于 2个第2小和第3小的数之和
// n-1个数(a1,a2 ... an-1)之和 必然小于 n个数(b1,b2 ... bn)之和 其中(a1 < a2 < .. < an-1 < b1 < b2 < ... < bn)
// 所以这个问题就变成n个数最多分成几组
// 当第1组数量为1 第2组数量为2 ... 第k组数量为k时，分组数量最大
//
// 时间复杂度 O(1)
// 空间复杂度 O(1)

func maximumGroups(grades []int) int {
	ans := int(math.Sqrt(2 * float64(len(grades))))
	if (1+ans)*ans > 2*len(grades) {
		return ans - 1
	}
	return ans
}
