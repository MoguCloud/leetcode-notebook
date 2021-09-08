// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 59.65 % of golang submissions (2.2 MB)

func minSwaps(s string) int {
	zeroOnEven := 0 // 0 在偶数位的个数
	zeroCount := 0  // 0 的个数
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeroCount += 1
			if (i+1)&1 == 0 {
				zeroOnEven += 1
			}
		}
	}
	if abs(zeroCount-(len(s)-zeroCount)) > 1 {
		// 0 1 的个数差 > 1 时，无法组成01相间的数
		return -1
	}
	if len(s)&1 == 0 {
		// s 长度偶数
		// 0 全部在奇数位 或 0 全部在偶数位
		return min(abs(zeroOnEven-len(s)/2), zeroOnEven)
	} else {
		// s 长度奇数
		if zeroCount > len(s)-zeroCount {
			// 0 更多
			// 1 需要出现在偶数位上 即 所有偶数位上的0都要换成1
			return zeroOnEven
		} else {
			// 1 更多
			// 0 需要出现在偶数位上 即 所有不在偶数位上的0都要变成0
			return zeroCount - zeroOnEven
		}
	}

}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}
