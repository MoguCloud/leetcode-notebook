// Your runtime beats 88 % of golang submissions (167 ms)
// Your memory usage beats 28 % of golang submissions (21.5 MB)
//
// BFS
// 从城市0为中心向两侧遍历，将从两侧指向中心(城市0)的改为从中心指向两侧，计算需要修改的数量
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func minReorder(n int, connections [][]int) int {
	// 记录各个城市，到该城市的城市已经从该城市出发去的城市
	neighbor := make([][2][]int, n)
	// 用于记录有没有遍历到该城市
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		neighbor[i] = [2][]int{[]int{}, []int{}}
	}
	for _, conn := range connections {
		from := conn[0]
		to := conn[1]
		neighbor[from][1] = append(neighbor[from][1], to)
		neighbor[to][0] = append(neighbor[to][0], from)
	}
	ret := 0
	q := []int{0}
	for len(q) > 0 {
		newQ := []int{}
		for _, city := range q {
			if visited[city] {
				continue
			}
			for _, to := range neighbor[city][1] {
				// 从中心出发的城市，需要改成前往中心
				if !visited[to] {
					newQ = append(newQ, to)
					ret += 1
				}
			}
			for _, from := range neighbor[city][0] {
				// 到中心的城市
				if !visited[from] {
					newQ = append(newQ, from)
				}
			}
			visited[city] = true
		}
		q = newQ
	}
	return ret
}
