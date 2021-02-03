// Your runtime beats 100 % of golang submissions (28 ms)
// Your memory usage beats 100 % of golang submissions (6.6 MB)
//
// 顺序遍历，遇到“&”后将元素存入一个临时数组，再次遇到“&”则将临时数组的元素都存入ret，并清空临时数组。如果遇到“;”，则判断临时数组是否组成需要转义的字符串，是则将转义完的符合存入ret，不是则将临时数组里的元素都存入ret，并清空临时数组。

func entityParser(text string) string {
	ret := []byte{}
	queue := []byte{}
	push := false
	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			push = true
		}
		if text[i] == '&' && push == true {
			for j := 0; j < len(queue); j++ {
				ret = append(ret, queue[j])
			}
			queue = []byte{}
		}
		if text[i] == ';' && push == true {
			push = false
			switch string(queue) {
			case "&quot":
				ret = append(ret, '"')
			case "&apos":
				ret = append(ret, '\'')
			case "&amp":
				ret = append(ret, '&')
			case "&gt":
				ret = append(ret, '>')
			case "&lt":
				ret = append(ret, '<')
			case "&frasl":
				ret = append(ret, '/')
			default:
				for j := 0; j < len(queue); j++ {
					ret = append(ret, queue[j])
				}
				ret = append(ret, ';')
			}
			queue = []byte{}
			continue
		}
		if push {
			queue = append(queue, text[i])
		} else {
			ret = append(ret, text[i])
		}
	}
	for j := 0; j < len(queue); j++ {
		ret = append(ret, queue[j])
	}
	return string(ret)
}
