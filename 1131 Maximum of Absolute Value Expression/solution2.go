// Math
//
// Your runtime beats 73.33 % of golang submissions (40 ms)
// Your memory usage beats 80 % of golang submissions (6.5 MB)
//
//       |xi - xj| + |yi - yj|
// = max((xi - xj + yi - yj ),  // 1
//       (xj - xi + yi - yj ),  // 2
//       (xi - xj + yj - yi ),  // 3
//       (xj - xi + yj - yi ))  // 4
// 因为i, j可以互换， 所以 1 = 4, 2 = 3
//       |xi - xj| + |yi - yj|
// = max((xi - xj + yi - yj ),
//       (xj - xi + yi - yj ))
//
//       |xi - xj| + |yi - yj| + |i - j|
// = max((xi - xj + yi - yj + i - j ),   = (xi + yi + i) - (xj + yj + j)
//       (xj - xi + yi - yj + i - j ),   = (-xi + yi + i) - (-xj + yj + j)
//       (xi - xj + yi - yj + j - i ),   = (xi + yi - i) - (xj + yj - j)
//       (xj - xi + yi - yj + j - i ),)  = (-xi + yi - i) - (-xj + yj - j)
//
// 令 A(i) = xi + yi + i
//    B(i) = -xi + yi + i
//    C(i) = xi + yi - i
//    D(i) = -xi + yi - i
//
// |xi - xj| + |yi - yj| + |i - j| = max( max(A) - min(A),
//                                        max(B) - min(B), 
//                                        max(C) - min(C), 
//                                        max(D) - min(D) )
// 
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	max1 := arr1[0] + arr2[0]
	min1 := arr1[0] + arr2[0]
	max2 := -arr1[0] + arr2[0]
	min2 := -arr1[0] + arr2[0]
	max3 := arr1[0] + arr2[0]
	min3 := arr1[0] + arr2[0]
	max4 := -arr1[0] + arr2[0]
	min4 := -arr1[0] + arr2[0]
	m := 0
	for i := 0; i < len(arr1); i++ {
		m = arr1[i] + arr2[i] + i
		max1 = max(max1, m)
		min1 = min(min1, m)
		m = -arr1[i] + arr2[i] + i
		max2 = max(max2, m)
		min2 = min(min2, m)
		m = arr1[i] + arr2[i] - i
		max3 = max(max3, m)
		min3 = min(min3, m)
		m = -arr1[i] + arr2[i] - i
		max4 = max(max4, m)
		min4 = min(min4, m)

	}
	return max(max1-min1, max2-min2, max3-min3, max4-min4)
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func max(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}

func min(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}
	return m
}
