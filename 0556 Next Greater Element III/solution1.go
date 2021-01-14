// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (1.9 MB)
//
// arr[i] 从i = 1 到 len(arr) 找到第一个 arr[i] < arr[i - 1]
// 再从 arr[j] 从 j = 0 到 i 找到第一个 arr[j] > arr[i]
// 交换 arr[i], arr[j] 再将 arr[:i] 从小到大排序
//
// 1 2 6 5 3 1
// 第一步找到 i = 4 即 arr[i] = 2
// 第二步找到 j = 3 即 arr[j] = 3
// 交换完之后变成 1 3 6 5 2 1
// 再将 1 3 (6 5 2 1)    括号内排序 即arr[0:4]进行排序 1 3 (1 2 5 6)

func nextGreaterElement(n int) int {
	// n按十进制位转换成数组
	arr := []int{}
	for n > 0 {
		arr = append(arr, n%10)
		n = n / 10
	}

	// 找到i
	i := 0
	find := false
	for i = 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			find = true
			break
		}
	}
	if !find {
		return -1
	}
	// 找到j
	j := 0
	for j = 0; j < i; j++ {
		if arr[j] > arr[i] {
			break
		}
	}
	// 交换
	arr[i], arr[j] = arr[j], arr[i]

	// 排序
	sort.Sort(sort.Reverse(sort.IntSlice(arr[:i])))

	// 判断溢出
	// 2^31 - 1 = 2147483647
	if len(arr) == 10 {
		if arr[9] > 2 {
			return -1
		}
		if arr[9] == 2 {
			num := 0
			for i := 8; i >= 0; i-- {
				num = num*10 + arr[i]
			}
			if num > 147483647 {
				return -1
			} else {
				return arr[9]*1000000000 + num
			}
		}
	}
	num := 0
	for i := len(arr) - 1; i >= 0; i-- {
		num = num*10 + arr[i]
	}
	return num
}
