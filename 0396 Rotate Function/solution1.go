// Your runtime beats 29.63 % of golang submissions (144 ms)
// Your memory usage beats 7.41 % of golang submissions (8.3 MB)
//
// 递推公式
// f(n) = f(n-1) - len(nums) * nums[len(nums) - n] + sum(nums)
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func maxRotateFunction(nums []int) int {
	prev := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		prev = prev + i*nums[i]
		sum = sum + nums[i]
	}
	max := prev
	for i := 1; i < len(nums); i++ {
		curr := prev - len(nums)*nums[len(nums)-i] + sum
		prev = curr
		if curr > max {
			max = curr
		}
	}
	return max
}
