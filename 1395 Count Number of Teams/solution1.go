// Your runtime beats 75.58 % of golang submissions (12 ms)
// Your memory usage beats 20.93 % of golang submissions (2.9 MB)
//
// biggerCounts[i] 为 rating 数组中索引大于i且值大于rating[i]的个数。
// 索引大于i，值小于rating[i]的个数就为len - 1 - i - biggerCounts[i]。
// 再次遍历 rating，当 rating[i] > rating[j] (i < j) 时，则能组成biggerCounts[j]个递减组合；
//                  当 rating[i] < rating[j] (i < j) 时，则能组成lessCounts[j]个递增组合。
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func numTeams(rating []int) int {
	biggerCounts := make([]int, len(rating))
	for i := 0; i < len(rating); i++ {
		count := 0
		for j := i + 1; j < len(rating); j++ {
			if rating[j] > rating[i] {
				count += 1
			}
		}
		biggerCounts[i] = count
	}
	ret := 0
	for i := 0; i < len(rating); i++ {
		for j := i + 1; j < len(rating); j++ {
			if rating[i] < rating[j] {
				ret += biggerCounts[j]
			} else {
				ret += len(rating) - 1 - j - biggerCounts[j]
			}
		}
	}
	return ret
}
