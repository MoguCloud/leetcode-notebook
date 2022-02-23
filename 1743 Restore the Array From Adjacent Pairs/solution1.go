// Your runtime beats 100 % of golang submissions (200 ms)
// Your memory usage beats 84 % of golang submissions (29 MB)
//
// 用哈希表记录每个元素的相邻元素
// 遍历哈希表，相邻元素只有1个的为起始或者末尾元素
// 再根据起始元素开始构造数组，后一项即为它的相邻元素且不等于前一项的元素
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func restoreArray(adjacentPairs [][]int) []int {
	if len(adjacentPairs) == 1 {
		return adjacentPairs[0]
	}
	neighborMap := make(map[int][]int)
	for _, pair := range adjacentPairs {
		neighborMap[pair[0]] = append(neighborMap[pair[0]], pair[1])
		neighborMap[pair[1]] = append(neighborMap[pair[1]], pair[0])
	}
	ret := make([]int, len(adjacentPairs)+1)
	var head int
	for k, v := range neighborMap {
		if len(v) == 1 {
			head = k
			break
		}
	}
	ret[0] = head
	prev := head
	next := neighborMap[head][0]
	for i := 1; i < len(ret); i++ {
		ret[i] = next
		if i == len(ret)-1 {
			break
		}
		neighbor := neighborMap[next]
		if neighbor[0] == prev {
			prev = next
			next = neighbor[1]
		} else {
			prev = next
			next = neighbor[0]
		}
	}
	return ret
}
