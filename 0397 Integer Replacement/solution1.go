// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (1.9 MB)
//
// 当 n 为偶数时，则除以 2
// 当 n 为奇数时，(n + 1) / 2 为偶数时，则 n = n + 1
//				 (n - 1) / 2 为偶数时，则 n = n - 1
//				  特殊情况 n == 3  时，则 n = n - 1

func integerReplacement(n int) int {
	ret := 0
	for n != 1 {
		if n&1 == 0 {
			n = n >> 1
		} else {
			if n == 3 {
				n = n - 1
			} else if ((n-1)/2)&1 == 0 {
				n = n - 1
			} else {
				n = n + 1
			}
		}
		ret += 1
	}
	return ret
}
