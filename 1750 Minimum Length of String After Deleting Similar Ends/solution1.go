// Your runtime beats 70 % of golang submissions (8 ms)
// Your memory usage beats 20 % of golang submissions (6.6 MB)
//
// 双指针
// 判断字符串开始和结束的元素是否相同，不相同的话则停止
// 相同的话则分别将头部和尾部连续相同的字符进行消去，即将指针移到下(上)一个不相同的元素
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func minimumLength(s string) int {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			break
		}
		ii := i + 1
		jj := j - 1
		for ii <= j {
			if s[ii] == s[i] {
				ii += 1
			} else {
				break
			}
		}
		for jj >= i {
			if s[jj] == s[j] {
				jj -= 1
			} else {
				break
			}
		}
		if ii > jj {

			return 0
		}
		i = ii
		j = jj
	}
	return j - i + 1
}
