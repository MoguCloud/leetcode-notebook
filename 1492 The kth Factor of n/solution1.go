// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 36.36 % of golang submissions (2 MB)
//
// prefix 升序存放 n 在 [0, sqrt(n)] 范围内的因子
// suffix 降序存放 n 在 (sqrt(n), n] 范围内的因子
// i = [0, sqrt(n)] 遍历，如果 n 能被 i 整除，则 i 存入 prefix, n / i 存入 suffix
// 如果 k <= len(prefix)，则第 k 个因子是 prefix 的第 k 项
// 如果 len(prefix) < k <= len(prefix) + len(suffix)，则第 k 个因子是 suffix 逆序的第 k - len(prefix) 项
// 如果 k > len(prefix) + len(suffix)，则 k 不在范围内
//
// 时间复杂度 O(sqrt(n))
// 空间复杂度 O(sqrt(n))

func kthFactor(n int, k int) int {
	prefix := []int{}
	suffix := []int{}
	for i := 1; i <= int(math.Sqrt(n)); i++ {
		if n%i == 0 {
			prefix = append(prefix, i)
			if i != n/i {
				suffix = append(suffix, n/i)
			}
			k = k - 1
			if k == 0 {
				return i
			}
		}
	}
	if k > len(suffix) {
		return -1
	}
	return suffix[len(suffix)-k]
}
