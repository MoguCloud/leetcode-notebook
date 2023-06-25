// Your runtime beats 64.29 % of golang submissions (1 ms)
// Your memory usage beats 100 % of golang submissions (1.9 MB)
//
// 逆向模拟
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func decodeAtIndex(s string, k int) string {
	size := 0
	for _, b := range s {
		if b >= '0' && b <= '9' {
			size *= int(b - '0')
		} else {
			size++
		}
	}
	var ret string
	for i := len(s) - 1; i >= 0; i-- {
		if k%size == 0 {
			for k := i; k >= 0; k-- {
				if !(s[k] >= '0' && s[k] <= '9') {
					return string(s[k])
				}
			}
		} else {
			k %= size
		}
		if s[i] >= '0' && s[i] <= '9' {
			size /= int(s[i] - '0')
		} else {
			size--
		}
	}
	return ret
}
