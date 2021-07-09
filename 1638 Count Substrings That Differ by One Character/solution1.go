// Your runtime beats 38.46 % of golang submissions (8 ms)
// Your memory usage beats 100 % of golang submissions (2 MB)
//
// Brute Force

func countSubstrings(s string, t string) int {
	ret := 0
	length := 0
	if len(s) > len(t) {
		length = len(t)
	} else {
		length = len(s)
	}
	for i := 1; i <= length; i++ {
		ret += check(&s, &t, i)
	}
	return ret
}

func check(s *string, t *string, length int) int {
	ret := 0
	for i := 0; i <= len(*s)-length; i++ {
		for j := 0; j <= len(*t)-length; j++ {
			diff := 0
			for k := 0; k < length; k++ {
				if (*s)[i+k] != (*t)[j+k] {
					diff += 1
				}
				if diff > 1 {
					break
				}
			}
			if diff == 1 {
				ret += 1
			}
		}
	}
	return ret
}
