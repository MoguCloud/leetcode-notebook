// Force Brute
// 
// Your runtime beats 6.67 % of golang submissions (3396 ms)
// Your memory usage beats 100 % of golang submissions (6.4 MB)

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	max := calc(&arr1, &arr2, 0, 0)
	for i := 0; i < len(arr1); i++ {
		for j := i; j < len(arr1); j++ {
			if calc(&arr1, &arr2, i, j) > max {
				max = calc(&arr1, &arr2, i, j)
			}
		}
	}
	return max
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func calc(arr1 *[]int, arr2 *[]int, i int, j int) int {
	return abs((*arr1)[i]-(*arr1)[j]) + abs((*arr2)[i]-(*arr2)[j]) + abs(i-j)
}
