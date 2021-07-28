// Your runtime beats 99.49 % of golang submissions (8 ms)
// Your memory usage beats 99.49 % of golang submissions (6.5 MB)
//
// 双指针
// 从左往右进行遍历
// 当前骨牌i为 '.' 时，从当前骨牌i往后遍历直到遇到一个不是 '.' 的骨牌j
//					  如果遍历完都没有遇到不是 '.' 的，从当前骨牌i到结束都是'.'
//					  如果骨牌j为'R'，则骨牌i到骨牌j都是'.'  (1)
//					  如果骨牌j为'L'，则骨牌i到骨牌j都是'L'  (2)
// 当前骨牌i为 'R' 时，从当前骨牌i往后遍历直到遇到一个不是 '.' 的骨牌j
//					  如果遍历完都没有遇到不是 '.' 的，从当前骨牌i到结束都是'R' (3)
//					  如果骨牌j为'R'，则骨牌i到骨牌j都是'R'  (4)
//                    如果骨牌j为'L'，则骨牌i往右变为'R'，骨牌j往左变为'L'(可以通过双指针实现)  (5)
// 当前骨牌i为 'L' 时，因为上一个节点必然只会是'L'，所以只要往下遍历即可
// 					  因为之前遍历的结果只会是上述情况(1)(2)(3)(4)之一，
//                    而(1)(3)是遍历结束，(2)(5)的上一个节点是'L'，而情况(4)会继续在(3)(4)(5)中直到变成(3)或(5)
//
//  时间复杂度 O(n)
//  空间复杂度 O(n)

func pushDominoes(dominoes string) string {
	byteArr := []byte(dominoes)
	i := 0
	for i < len(byteArr) {
		if byteArr[i] == '.' {
			findNext := false
			for j := i + 1; j < len(byteArr); j++ {
				if byteArr[j] == 'R' {
					findNext = true
					i = j
					break
				} else if byteArr[j] == 'L' {
					findNext = true
					for k := i; k <= j; k++ {
						byteArr[k] = 'L'
					}
					i = j
					break
				}
			}
			if !findNext {
				break
			}
		} else if byteArr[i] == 'L' {
			i += 1
		} else {
			findNext := false
			for j := i + 1; j < len(byteArr); j++ {
				if byteArr[j] == 'R' {
					findNext = true
					for k := i; k <= j; k++ {
						byteArr[k] = 'R'
					}
					i = j
					break
				} else if byteArr[j] == 'L' {
					findNext = true
					for ii, jj := i, j; ii < jj; ii, jj = ii+1, jj-1 {
						byteArr[ii] = 'R'
						byteArr[jj] = 'L'
					}
					i = j
					break
				}
			}
			if !findNext {
				for i < len(byteArr) {
					byteArr[i] = 'R'
					i += 1
				}
			}
		}
	}
	return string(byteArr)
}
