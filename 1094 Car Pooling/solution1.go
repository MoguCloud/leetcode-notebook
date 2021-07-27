// Your runtime beats 94.87 % of golang submissions(4 ms)
// Your memory usage beats 12.82 % of golang submissions (4.3 MB)
//
// 哈希表 upMap key: 上车地点 value: 上车人数
// 哈希表 downMap key: 下车地点 value: 下车人数
// 数组 stops 上下车的地点，升序排序
// 按 stops 进行遍历，即按上下车地点顺序遍历
// 到站的时候计算该站 减去下车人数，加上上车人数
// 若 > capacity，则无法拼车；遍历完都 <= capacity，则可以拼车
//
// 时间复杂度 O(n*logn)
// 空间复杂度 O(n)

func carPooling(trips [][]int, capacity int) bool {
	if len(trips) == 0 {
		return true
	}

	stopsMap := make(map[int]struct{})
	upMap := make(map[int]int)
	downMap := make(map[int]int)
	for i := 0; i < len(trips); i++ {
		stopsMap[trips[i][1]] = struct{}{}
		stopsMap[trips[i][2]] = struct{}{}
		upMap[trips[i][1]] += trips[i][0]
		downMap[trips[i][2]] += trips[i][0]
	}
	stops := []int{}
	for k, _ := range stopsMap {
		stops = append(stops, k)
	}
	sort.Ints(stops)

	passengers := upMap[stops[0]]

	if passengers > capacity {
		return false
	}

	for i := 1; i < len(stops); i++ {
		passengers -= downMap[stops[i]]
		passengers += upMap[stops[i]]
		if passengers > capacity {
			return false
		}
	}
	return true
}
