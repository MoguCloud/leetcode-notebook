// Your runtime beats 50 % of golang submissions (793 ms)
// Your memory usage beats 83.33 % of golang submissions (12 MB)
//
// 利用 map 判断是否是子集

func peopleIndexes(favoriteCompanies [][]string) []int {
	ret := []int{}
	for i := 0; i < len(favoriteCompanies); i++ {
		isSub := false
		for j := 0; j < len(favoriteCompanies); j++ {
			if i == j {
				continue
			}
			if isSubSet(favoriteCompanies[j], favoriteCompanies[i]) {
				isSub = true
				break
			}
		}
		if !isSub {
			ret = append(ret, i)
		}
	}
	return ret
}

func isSubSet(sliceA, sliceB []string) bool {
	// sliceB 是否是 sliceA 的子集
	if len(sliceB) >= len(sliceA) {
		return false
	}
	setA := make(map[string]struct{})
	for _, strA := range sliceA {
		setA[strA] = struct{}{}
	}
	for _, strB := range sliceB {
		if _, ok := setA[strB]; !ok {
			return false
		}
	}
	return true
}
