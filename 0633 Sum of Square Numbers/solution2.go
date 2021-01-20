// 双指针
//
// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 29.63 % of golang submissions (1.9 MB)
//
// 时间复杂度 O(sqrt(n))
// 空间复杂度 O(1)

func judgeSquareSum(c int) bool {
	i := 0
	j := int(math.Sqrt(float64(c)))
	ret := false
	for i <= j {
		sum := i * i + j * j
		if sum == c {
			ret = true
			break
		}
		if sum > c {
			j -= 1
		}
		if sum < c {
			i += 1
		}
	}
	return ret
}
