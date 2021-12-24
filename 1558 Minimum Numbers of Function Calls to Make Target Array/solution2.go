// Your runtime beats 75 % of golang submissions (41 ms)
// Your memory usage beats 31.25 % of golang submissions (7.5 MB)
//
// 给定函数可以实现
// 1. 数组中的1个数 +1
// 2. 数组中的全部数都 x2
// 倒着推就是可以将某个数-1或者所有数都除以2
// 要使函数调用次数最少，应该尽可能使用x2
// 即倒推情况下，如果没有奇数的情况下则使用/2，出现奇数时将奇数-1变成偶数
// 因为奇数-1只能1次1次调用函数实现，偶数/2可以调用1次函数实现
// 所以调用/2的次数为数组中最大数调用/2的次数
// 调用函数实现-1可以看作为二进制中最低位的1变成0
// 调用函数实现/2可以看作为二进制中的右移1位
// 所以每个数需要调用-1的次数为二进制中1的个数

func minOperations(nums []int) int {
	ret := 0
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		// 判断1的个数
		for num > 0 {
			if num&1 == 1 {
				ret += 1
			}
			num >>= 1
		}
	}
	// 判断需要/2的次数
	for maxNum > 1 {
		maxNum >>= 1
		ret += 1
	}
	return ret
}
