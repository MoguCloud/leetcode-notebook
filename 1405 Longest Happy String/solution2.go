// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 100 % of golang submissions (2 MB)
//
// 贪心
//
// 每次都选择剩余个数最多的字母，当末尾有连续2个剩余个数最多的字母时，选择剩余个数次多的字母
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

var ret []byte

func longestDiverseString(a int, b int, c int) string {
	ret = []byte{}
	for {
		flag := true
		flag, a, b, c = insert(a, b, c, 'a', 'b', 'c')
		if !flag {
			flag, b, a, c = insert(b, a, c, 'b', 'a', 'c')
		}
		if !flag {
			flag, c, a, b = insert(c, a, b, 'c', 'a', 'b')
		}
		if !flag {
			break
		}
	}
	return string(ret)
}

func insert(a int, b int, c int, charA byte, charB byte, charC byte) (bool, int, int, int) {
	if a >= b && a >= c && a != 0 {
		if len(ret) >= 2 && ret[len(ret)-1] == charA && ret[len(ret)-2] == charA {
			if b >= c && b != 0 {
				ret = append(ret, charB)
				return true, a, b - 1, c
			} else if c >= b && c != 0 {
				ret = append(ret, charC)
				return true, a, b, c - 1
			}
			return false, a, b, c
		} else {
			ret = append(ret, charA)
			return true, a - 1, b, c
		}
	} else {
		return false, a, b, c
	}
}
