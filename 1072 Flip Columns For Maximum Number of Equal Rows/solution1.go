// Your runtime beats 80 % of golang submissions (184 ms)
// Your memory usage beats 40 % of golang submissions (7.8 MB)
//
// 翻转后能够相等的行都是 所有数相同 或者所有数都不同的
// 所以只要统计 所有数都相同或者所有数都不同的 最大个数即可
// 使用哈希表分组统计，取最大值
// 为了将所有相同以及所有都不同分在一个组里，将第1位为1的都进行取反，即能分到同一组里
//
// 时间复杂度 O(M * N)
// 空间复杂度 O(M * N)

func maxEqualRowsAfterFlips(matrix [][]int) int {
	ret := 0
	memo := make(map[string]int)
	for i := 0; i < len(matrix); i++ {
		rowToBytes := []byte{}
		reverse := 0 // 通过异或来实现翻转
		if matrix[i][0] == 1 {
			reverse = 1
		}
		for j := 0; j < len(matrix[i]); j++ {
			rowToBytes = append(rowToBytes, byte(matrix[i][j]^reverse))
		}
		rowToStr := string(rowToBytes)
		count, _ := memo[rowToStr]
		count += 1
		memo[rowToStr] = count
		if count > ret {
			ret = count
		}
	}
	return ret
}
