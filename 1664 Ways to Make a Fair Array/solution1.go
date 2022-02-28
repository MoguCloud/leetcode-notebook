// Your runtime beats 83.33 % of golang submissions (147 ms)
// Your memory usage beats 38.89 % of golang submissions (10.7 MB)
//
// 前缀和
// 先遍历一遍，分别求出奇数位的前缀和以及偶数位的前缀和
// i从0到-1再遍历一遍，将i删除后，i之前的奇数和为之前的奇数和，i之后的奇数和为之前的偶数和
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func waysToMakeFair(nums []int) int {
	oddPrefixSum := []int{0}
	evenPrefixSum := []int{0}
	oddSum := 0
	evenSum := 0
	for i := 0; i < len(nums); i++ {
		if (i+1)&1 == 1 {
			oddSum += nums[i]
		} else {
			evenSum += nums[i]
		}
		oddPrefixSum = append(oddPrefixSum, oddSum)
		evenPrefixSum = append(evenPrefixSum, evenSum)
	}
	ret := 0
	for i := 0; i < len(nums); i++ {
		oddSum := 0
		evenSum := 0
		oddSum += oddPrefixSum[i] - oddPrefixSum[0]
		evenSum += evenPrefixSum[i] - evenPrefixSum[0]
		oddSum += evenPrefixSum[len(evenPrefixSum)-1] - evenPrefixSum[i]
		evenSum += oddPrefixSum[len(oddPrefixSum)-1] - oddPrefixSum[i]
		if (i+1)&1 == 0 {
			oddSum -= nums[i]
		} else {
			evenSum -= nums[i]
		}
		if oddSum == evenSum {
			ret += 1
		}
	}
	return ret
}
