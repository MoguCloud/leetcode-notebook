// Your runtime beats 77.78 % of golang submissions (37 ms)
// Your memory usage beats 66.67 % of golang submissions (2.5 MB)
//
// 枚举排列
//
// 时间复杂度 O(m*m*n+m*m!)
// 空间复杂度 O(m*m)

var maxCap int

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
	m := len(students)
	// 先预处理各个学生和导师的兼容性
	memo := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			memo[i*10+j] = getCompatibility(students[i], mentors[j])
		}
	}
	// 枚举所有学生和导师组合的兼容性评分和，然后取最大值
	maxCap = 0
	dfs(0, 0, m, 0, &memo)
	return maxCap
}

func dfs(studentIdx int, usedMentor int, m int, cap int, memo *map[int]int) {
	if studentIdx >= m {
		if cap > maxCap {
			maxCap = cap
		}
		return
	}
	// 使用二进制来代替布尔数组，用来记录导师是否已经和学生配对过
	for i := 0; i < m; i++ {
		if usedMentor&(1<<i) == 0 {
			dfs(studentIdx+1, usedMentor|(1<<i), m, cap+(*memo)[studentIdx*10+i], memo)
		}
	}
	return
}

func getCompatibility(s []int, m []int) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		if s[i] == m[i] {
			ret += 1
		}
	}
	return ret
}
