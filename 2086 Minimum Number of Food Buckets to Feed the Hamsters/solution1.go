// Your runtime beats 100 % of golang submissions (3 ms)
// Your memory usage beats 40 % of golang submissions (6.7 MB)
//
// 从左往右遍历
// 如果房间两侧有水桶了，则不需要放置水桶
// 如果没有水桶，优先放置在右侧，右侧不能放置时，才放置在左侧
//   因为放在右侧的水桶可以被右边的房间也利用到。
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func minimumBuckets(hamsters string) int {
	ret := 0
	arr := []byte(hamsters)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 'H' {
			if i-1 >= 0 && arr[i-1] == 'B' {
				// pass
			} else {
				place := false
				// 尝试放在后面
				if i+1 < len(hamsters) {
					if arr[i+1] == '.' {
						arr[i+1] = 'B'
						ret += 1
						place = true
					}
				}
				// 无法放在后面，需要放在前面
				if !place {
					if i-1 >= 0 {
						if arr[i-1] == '.' {
							arr[i-1] = 'B'
							ret += 1
							place = true
						}
					}
				}
				// 前后都无法放置
				if !place {
					return -1
				}
			}
		}
	}
	return ret
}
