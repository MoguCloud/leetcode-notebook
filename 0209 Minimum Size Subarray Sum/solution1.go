// Your runtime beats 49.03 % of golang submissions (10 ms)
// Your memory usage beats 83.23 % of golang submissions (3.8 MB)
//
// 滑动窗口
// 用i和j分别表示左边和右边的指针
// 当滑动窗口中的和sum > target，左指针右移，并判断长度是否最短
// 当 <= taget，右指针右移
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func minSubArrayLen(target int, nums []int) int {
	sum := nums[0]
	ret := len(nums) + 1
	i := 0
	j := 0
	for {
		if sum >= target {
			count := j - i + 1
			if count < ret {
				ret = count
			}
		}
		if sum >= target {
			sum -= nums[i]
			i += 1
			if i >= len(nums) {
				break
			}
		} else {
			j += 1
			if j >= len(nums) {
				break
			}
			sum += nums[j]
		}
	}
	if ret == len(nums)+1 {
		ret = 0
	}
	return ret
}
