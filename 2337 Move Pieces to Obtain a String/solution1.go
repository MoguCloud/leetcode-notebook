// Your runtime beats 14.29 % of golang submissions (1450 ms)
// Your memory usage beats 100 % of golang submissions (6.80 MB)
//
// 模拟
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func canChange(start string, target string) bool {
	startSlice := []byte(start)
	for i := 0; i < len(start); i++ {
		if startSlice[i] == target[i] {
			continue
		}
		if target[i] == 'R' && startSlice[i] == 'L' {
			return false
		} else if target[i] == 'R' && startSlice[i] == '_' {
			return false
		} else if target[i] == 'L' && startSlice[i] == 'R' {
			return false
		} else if target[i] == 'L' && startSlice[i] == '_' {
			find := false
			for j := i + 1; j < len(startSlice); j++ {
				if startSlice[j] == 'R' {
					break
				}
				if startSlice[j] == 'L' {
					find = true
					startSlice[i], startSlice[j] = startSlice[j], startSlice[i]
					break
				}
			}
			if !find {
				return false
			}
		} else if target[i] == '_' && startSlice[i] == 'L' {
			return false
		} else if target[i] == '_' && startSlice[i] == 'R' {
			find := false
			for j := i + 1; j < len(startSlice); j++ {
				if startSlice[j] == 'L' {
					break
				}
				if startSlice[j] == '_' {
					find = true
					startSlice[i], startSlice[j] = startSlice[j], startSlice[i]
					break
				}
			}
			if !find {
				return false
			}
		}
	}
	return true
}
