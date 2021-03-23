// 模拟对角线遍历
//
// Time Limit Exceeded

func findDiagonalOrder(nums [][]int) []int {
	rowMap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		rowMap[i] = struct{}{}
	}
	rows := len(nums)
	i := 0
	ret := []int{}
	for rows > 0 {
		for j := len(nums) - 1; j >= 0; j-- {
			if i-j < 0 {
				continue
			}
			if _, ok := rowMap[j]; !ok {
				continue
			}
			if i-j >= len(nums[j]) {
				rows -= 1
				delete(rowMap, j)
				continue
			}
			ret = append(ret, nums[j][i-j])
		}
		i += 1
	}
	return ret
}
