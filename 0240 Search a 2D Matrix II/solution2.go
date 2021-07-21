// Your runtime beats 72.00 % of golang submissions (24 ms)
// Your memory usage beats 49.14 % of golang submissions (6.9 MB)
//
// 从左下角开始遍历
// 如果当前数字 > target，则 target 不在当前行，因为当前数字是当前行最小的，当前行都比target要大，所以往上移动一行(逻辑上删除当前行)
// 如果当前数字 < target，则 target 不在当前列，因为当前数字是当前列最大的，当前列都比target要小，所以往右移动一行(逻辑上删除当前列)
//
// 时间复杂度 O(m+n)
// 空间复杂度 O(1)

func searchMatrix(matrix [][]int, target int) bool {
	i := 0
	j := len(matrix[0]) - 1
	for i >= 0 && i <= len(matrix)-1 && j >= 0 && j <= len(matrix[0])-1 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j -= 1
		} else {
			i += 1
		}
	}
	return false
}
