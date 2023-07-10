// Your runtime beats 90.91 % of golang submissions (67 ms)
// Your memory usage beats 100 % of golang submissions (7.3 MB)
//
// 模拟
//
// 时间复杂度 O(len(nums) + len(queries))
// 空间复杂度 O(n)

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	ans := make([]int, len(nums))
	evenSum := 0
	先计算偶数和
	for _, num := range nums {
		if num&1 == 0 {
			evenSum += num
		}
	}
	// 遍历每次查询
	for i, query := range queries {
		val := query[0]
		idx := query[1]
		// 修改前是否为偶数
		if nums[idx]&1 == 0 {
			evenSum -= nums[idx]
		}
		// 进行修改
		nums[idx] += val
		// 修改后是否为偶数
		if nums[idx]&1 == 0 {
			evenSum += nums[idx]
		}
		ans[i] = evenSum
	}
	return ans
}
