// 贪心
//
// Your runtime beats 73.78 % of golang submissions (92 ms)
// Your memory usage beats 87.81 % of golang submissions (6.9 MB)
//
// 体重最大的和体重最小的之和 <= limit，则最重和最轻的一起，否在最重的自己一艘

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i := 0
	j := len(people) - 1
	ret := 0
	for i < j {
		if people[j]+people[i] <= limit {
			ret += 1
			i += 1
			j -= 1
		} else {
			ret += 1
			j -= 1
		}
	}
	if i == j {
		ret += 1
	}
	return ret
}
