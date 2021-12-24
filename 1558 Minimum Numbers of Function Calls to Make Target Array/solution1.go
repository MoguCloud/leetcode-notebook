// Time Limit Exceeded

func minOperations(nums []int) int {
	ret := 0
	for {
		// 检查是否全部都变成0了，全变成0则退出
		finished := true
		for _, num := range nums {
			if num != 0 {
				finished = false
				break
			}
		}
		if finished {
			break
		}
		// 如果存在奇数，则让奇数-1
		has_odd := false
		for idx, num := range nums {
			if num&1 == 1 {
				ret += 1
				nums[idx] = num - 1
				has_odd = true
				break
			}
		}
		if has_odd {
			continue
		}
		// 如果不存在奇数，则让偶数/2
		for idx, num := range nums {
			nums[idx] = num / 2
		}
		ret += 1
	}
	return ret
}
