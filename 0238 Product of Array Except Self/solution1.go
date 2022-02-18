// Your runtime beats 66.8 % of golang submissions (29 ms)
// Your memory usage beats 33.2 % of golang submissions (8.3 MB)
//
// prefixProduct 前缀积数组 prefixProduct[i] = nums[0] * nums[1] * .. * nums[i]
// suffixProduct 后缀积数组 suffixProduct[i] = nums[i] * nums[i+1] * .. * nums[len(nums)-1]
// 返回结果ret数组 ret[i] = prefixProduct[i-1] * suffixProduct[i+1]
// 当 i = 0 或者 len(nums) - 1时进行特殊处理即可
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func productExceptSelf(nums []int) []int {
	prefixProduct := make([]int, len(nums))
	suffixProduct := make([]int, len(nums))
	prefixProduct[0] = nums[0]
	suffixProduct[len(suffixProduct)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		prefixProduct[i] = prefixProduct[i-1] * nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		suffixProduct[i] = suffixProduct[i+1] * nums[i]
	}
	ret := make([]int, len(nums))
	ret[0] = suffixProduct[1]
	ret[len(ret)-1] = prefixProduct[len(ret)-2]
	for i := 1; i < len(nums)-1; i++ {
		ret[i] = prefixProduct[i-1] * suffixProduct[i+1]
	}
	return ret
}
