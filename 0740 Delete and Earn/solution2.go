// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 58.78 % of golang submissions (3.8 MB)
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
// 因为 dp[0][i] 和 dp[1][i] 只与 dp[0][i-1] 和 dp[1][i-1] 有关
// 所以可以用 2 个变量进行空间压缩
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

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
	dp0 := 0
	dp1 := noDupNums[0] * numCount[noDupNums[0]]
	ret := dp1
	for i := 1; i < len(noDupNums); i++ {
		newDp0 := max(dp0, dp1)
		newDp1 := 0
		if noDupNums[i] == noDupNums[i-1]+1 {
			newDp1 = dp0 + noDupNums[i]*numCount[noDupNums[i]]
		} else {
			newDp1 = max(dp0, dp1) + noDupNums[i]*numCount[noDupNums[i]]
		}
		if newDp1 > ret {
			ret = newDp1
		}
		dp0, dp1 = newDp0, newDp1
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
