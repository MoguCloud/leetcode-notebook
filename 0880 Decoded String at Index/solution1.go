// Runtime Error
//
// 正向展开

func decodeAtIndex(s string, k int) string {
	arr := []string{}
	length := 0
	var ret string
	for _, b := range s {
		if b >= '2' && b <= '9' {
			newStr := strings.Join(arr, "")
			count := int(b - '0')
			for i := 1; i < count; i++ {
				arr = append(arr, newStr)
			}
			length += len(newStr) * (count - 1)
		} else {
			str := string(b)
			arr = append(arr, str)
			length++
		}
		if k <= length {
			for _, str := range arr {
				if k <= len(str) {
					return string(str[k-1])
				}
				k -= len(str)
			}
		}
	}
	return ret
}
