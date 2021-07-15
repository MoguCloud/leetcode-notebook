// Your runtime beats 95.95 % of golang submissions (4 ms)
// Your memory usage beats 49.32 % of golang submissions (3.9 MB)
//
// DP
//
// 将数组进行去重排序后，即可转换成打家劫舍问题
// 相邻的数组如果相差1，则不能同时选
// noDupNums 为去重并且排序过的数组
// numCount key为 nums中的元素 value为该元素出现过得次数 的 哈希表
// dp为2维数组
// dp[0][i]为不选(删)该元素，最大的点数
// dp[0][i]为选(删)该元素，最大的点数
// 状态转移方程
// dp[0][i] = max(dp[0][i-1], dp[1][i-1])
// 当noDupNums[i]与当noDupNums[i-1]相差不为1时
// dp[1][i] = max(dp[0][i-1], dp[1][i-1]) + noDupNums[i] * numCount[i]
// 当noDupNums[i]与当noDupNums[i-1]相差为1时
// dp[1][i] = dp[0][i-1] + noDupNums[i] * numCount[i]
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func deleteAndEarn(nums []int) int {
	numCount := make(map[int]int)
	noDupNums := []int{}
	for i := 0; i < len(nums); i++ {
		_, ok := numCount[nums[i]]
		if !ok {
			noDupNums = append(noDupNums, nums[i])
		}
		numCount[nums[i]] += 1
	}
	sort.Ints(noDupNums)
	dp := make([][]int, 2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(noDupNums))
	}
	dp[0][0] = 0
	dp[1][0] = noDupNums[0] * numCount[noDupNums[0]]
	ret := dp[1][0]
	for i := 1; i < len(noDupNums); i++ {
		dp[0][i] = max(dp[0][i-1], dp[1][i-1])
		if noDupNums[i] == noDupNums[i-1]+1 {
			dp[1][i] = dp[0][i-1] + noDupNums[i]*numCount[noDupNums[i]]
		} else {
			dp[1][i] = max(dp[0][i-1], dp[1][i-1]) + noDupNums[i]*numCount[noDupNums[i]]
		}
		if dp[1][i] > ret {
			ret = dp[1][i]
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
