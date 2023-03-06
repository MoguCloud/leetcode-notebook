// Your runtime beats 100 % of golang submissions (3 ms)
// Your memory usage beats 80 % of golang submissions (5.5 MB)
// 
// 从左往右遍历
// 如果房间两侧有水桶了，则不需要放置水桶
// 如果没有水桶，优先放置在右侧，右侧不能放置时，才放置在左侧
//    因为放在右侧的水桶可以被右边的房间也利用到。
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func minimumBuckets(hamsters string) int {
	ret := 0
	i := 0
	for i < len(hamsters) {
		if hamsters[i] == 'H' {
			place := false
			// 尝试放在后面
			if i+1 < len(hamsters) {
				if hamsters[i+1] == '.' {
					ret += 1
					place = true
					// 如果放在后面成功的话，则第i+2个也不用判断，下次直接判断第i+3个
					i += 3
				}
			}
			// 无法放在后面，需要放在前面
			if !place {
				if i-1 >= 0 {
					if hamsters[i-1] == '.' {
						ret += 1
						place = true
						i += 1
					}
				}
			}
			// 前后都无法放置
			if !place {
				return -1
			}
		} else {
			i += 1
		}
	}
	return ret
}
