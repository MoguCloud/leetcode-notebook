// DFS
//
// Time Limit Exceeded

var maxDay int
var maxCount int
var max int

func maxEvents(events [][]int) int {
	maxCount = len(events)
	used := make(map[int]struct{})
	maxDay = 0
	max = 0
	for i := 0; i < len(events); i++ {
		if events[i][1] > maxDay {
			maxDay = events[i][1]
		}
	}
	dfs(&events, &used, 1, 0)
	return max
}

func dfs(events *[][]int, used *map[int]struct{}, day int, count int) int {
	result, ok := memo[mKey]
	if ok {
		return result
	}
	if max == maxCount {
		return count
	}
	if day > maxDay {
		return count
	}
	// 记录当前的最大值
	m := count
	// 当天参加了
	for i := 0; i < len(*events); i++ {
		// 判断此event是否参加过
		_, ok := (*used)[i]
		if !ok {
			// 判断时候在时间内
			if (*events)[i][0] <= day && day <= (*events)[i][1] {
				(*used)[i] = struct{}{}
				c := dfs(events, used, day+1, count+1)
				if c > max {
					max = c
				}
				if c > m {
					m = c
				}
				delete(*used, i)
			}
		}
	}
	// 当天未参加
	c := dfs(events, used, day+1, count)
	if m > c {
		m = c
	}
	return m
}
