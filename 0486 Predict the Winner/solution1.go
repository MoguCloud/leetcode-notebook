// Your runtime beats 25.00 % of golang submissions (4 ms)
// Your memory usage beats 12.50 % of golang submissions (2.8 MB)
//
// 递归
//
// 时间复杂度 O(2^n)
// 空间复杂度 O(n^2)

var memo map[string]int

func PredictTheWinner(nums []int) bool {
	// 用于记忆化
	memo = make(map[string]int)
	// 判断先手得分是否 >= 后手得分
	return helper(nums, 0, len(nums)-1, 1) >= 0
}

// helper 从i开始到j结束，先手turn为+1，后手turn为-1，先手记正分，后手记负分
func helper(nums []int, i int, j int, turn int) int {
	mapKey := fmt.Sprintf("%d %d %d", i, j, turn)
	// 判断是否命中缓存
	if v, ok := memo[mapKey]; ok {
		return v
	}
	// 当 i = j时，只有一个可选项
	if i == j {
		return turn * nums[i]
	}
	// 分别计算选择最前面和最后的值
	selectFirst := turn*nums[i] + helper(nums, i+1, j, -turn)
	selectLast := turn*nums[j] + helper(nums, i, j-1, -turn)
	// 选择最大值(后手是 - max(abs(result)) )
	ret := turn * max(turn*selectFirst, turn*selectLast)
	// 缓存结果
	memo[mapKey] = ret
	return ret

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
