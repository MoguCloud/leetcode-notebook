// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 16.51 % of golang submissions (2 MB)
//
// 递归
// x ** 5 = x * (x ** 4) = x * ((x * x) ** 2) = ...
// pow(x, n) = x * pow(x, n-1) 当 n 为奇数时
// pow(x, n) = pow(x*x, n/2)   当 n 为偶数时
//
// 时间复杂度 O(logn)
// 空间复杂度 O(logn)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n == 1 {
		return x
	}
	if n&1 == 1 {
		return x * myPow(x, n-1)
	} else {
		return myPow(x*x, n/2)
	}
}
