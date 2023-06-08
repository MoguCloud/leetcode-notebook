// Your runtime beats 52.17 % of golang submissions (52 ms)
// Your memory usage beats 73.91 % of golang submissions (7.2 MB)
//
// BFS
//
// 时间复杂度 O(n^3)
// 空间复杂度 O(n)

func canBomb(bombs [][]int, i int, j int) bool {
	d1 := (bombs[i][0]-bombs[j][0])*(bombs[i][0]-bombs[j][0]) + (bombs[i][1]-bombs[j][1])*(bombs[i][1]-bombs[j][1])
	return d1 <= bombs[i][2]*bombs[i][2]
}

func maximumDetonation(bombs [][]int) int {
	maxBomb := 0
	for i := 0; i < len(bombs); i++ {
		bombed := make([]bool, len(bombs))
		bombed[i] = true
		q := []int{i}
		count := 1
		for len(q) > 0 {
			newQ := []int{}
			for _, b := range q {
				for i := 0; i < len(bombs); i++ {
					if !bombed[i] && canBomb(bombs, b, i) {
						newQ = append(newQ, i)
						bombed[i] = true
						count++
					}
				}
			}
			q = newQ
		}
		if count > maxBomb {
			maxBomb = count
		}
	}
	return maxBomb
}
