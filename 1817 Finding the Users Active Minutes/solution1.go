// Your runtime beats 100 % of golang submissions (136 ms)
// Your memory usage beats 16.67 % of golang submissions (10.3 MB)
//
// uamMap 记录每个用户的uam
// userMinutes 记录每个用户在每一分钟是否操作过，便于计算uam
// 遍历uamMap 计算出每个uam对应的用户数
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	uamMap := make(map[int]int)
	userMinutes := make(map[int]map[int]bool)
	for _, log := range logs {
		id, minute := log[0], log[1]
		minuteMap, ok := userMinutes[id]
		if !ok {
			minuteMap = make(map[int]bool)
		}
		_, ok = minuteMap[minute]
		if !ok {
			minuteMap[minute] = true
			userMinutes[id] = minuteMap
			uamMap[id] += 1
		}
	}
	ret := make([]int, k)
	for _, uam := range uamMap {
		ret[uam-1] += 1
	}
	return ret
}

