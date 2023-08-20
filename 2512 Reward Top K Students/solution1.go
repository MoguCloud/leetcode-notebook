// Your runtime beats 63.16 % of golang submissions (106 ms)
// Your memory usage beats 26.32 % of golang submissions (8.8 MB)
//
// Hash + 排序

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	// 将数组转换成Hash方便比对
	posMap := make(map[string]struct{})
	negMap := make(map[string]struct{})
	for _, str := range positive_feedback {
		posMap[str] = struct{}{}
	}
	for _, str := range negative_feedback {
		negMap[str] = struct{}{}
	}
	stuArr := [][2]int{}
	for i := 0; i < len(report); i++ {
		// 分别计算每个学生的分数
		pt := 0
		strs := strings.Split(report[i], " ")
		for _, str := range strs {
			if _, ok := posMap[str]; ok {
				pt += 3
			}
			if _, ok := negMap[str]; ok {
				pt -= 1
			}
		}
		stuArr = append(stuArr, [2]int{student_id[i], pt})
	}
	// 按分数进行排序
	sort.Slice(stuArr, func(i, j int) bool {
		if stuArr[i][1] == stuArr[j][1] {
			return stuArr[i][0] < stuArr[j][0]
		}
		return stuArr[i][1] > stuArr[j][1]
	})
	// 取前K个
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = stuArr[i][0]
	}
	return ret
}
