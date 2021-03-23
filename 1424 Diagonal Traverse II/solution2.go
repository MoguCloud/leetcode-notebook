// Your runtime beats 42.42 % of golang submissions (244 ms)
// Your memory usage beats 15.15 % of golang submissions (25.1 MB)
//
// 二位数组里所有的数字按行号+列号升序排列即为答案。
// 当行号+列号相同时，则行号越大排在越前面。
// 将行号+列号相同的数字存在同一个数组里，按列号降序遍历即可保证相同时，行号大的在前面。
// 最后只需要按行号+列号升序将对应的数组进行合并即可。

func findDiagonalOrder(nums [][]int) []int {
	idxArr := []int{}
	numMap := make(map[int][]int) // key 行号+列号 value 值数组
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < len(nums[i]); j++ {
			arr, ok := numMap[i+j]
			if !ok {
				arr = []int{}
				idxArr = append(idxArr, i+j)
			}
			arr = append(arr, nums[i][j])
			numMap[i+j] = arr
		}
	}
	ret := []int{}
	sort.Ints(idxArr)
	for i := 0; i < len(idxArr); i++ {
		ret = append(ret, numMap[idxArr[i]]...)
	}
	return ret
}
