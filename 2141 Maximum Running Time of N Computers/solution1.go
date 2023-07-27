// Your runtime beats 100 % of golang submissions (109 ms)
// Your memory usage beats 100 % of golang submissions (8.99 MB)
//
// 二分搜索
// n 台电脑运行 t 时间总共需要电量n * t
// 一个电池在 t 时间内能够提供的电量为 batteries[i]，但是不能超过 t，因为不能同时给多个电脑使用，所以能提供的电量为 min(batteries[i], t)
// 所以只要找到 sum(min(batteries[i], t)) >= n * t，满足条件的t的最大值即可
// 搜索下界为1 搜索上界为1e14
// 具体证明参考 https://leetcode.cn/problems/maximum-running-time-of-n-computers/solution/er-fen-da-an-by-newhar-swi2/
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(1)

func maxRunTime(n int, batteries []int) int64 {
	var lo int64 = 0
	var hi int64 = int64(1e14)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if check(batteries, n, mid) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo - 1
}

// check 判断是否可以n个电脑使用t时间
func check(batteries []int, n int, t int64) bool {
	var sum int64 = 0
	for _, b := range batteries {
		// 对于一节电池而言，最多使用t时间，因为同一时间不能被使用多次
		sum += min(int64(b), t)
	}
	return sum >= int64(n)*t
}

func min(i, j int64) int64 {
	if i < j {
		return i
	}
	return j
}
