// Your runtime beats 100 % of golang submissions (117 ms)
// Your memory usage beats 100 % of golang submissions (9.4 MB)
//
// 一对可以被k整除的数可以是2个都被k整除的数，可以是一个被k除，余数是i,另一个被k除，余数是k-i的。只要2数被k除的余数之和为0或者k即可。
// 遍历数组arr，分别计算除以k的余数。使用长度为k的数组记录各余数出现的次数。
// 再分组统计是否各余数可以配对。
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func canArrange(arr []int, k int) bool {
	mods := make([]int, k)
	for _, n := range arr {
		if n%k < 0 {
			// 由于负数的余数会是负数，需要加一个k使得余数变成正数
			mods[(n%k)+k]++
		} else {
			mods[n%k]++
		}
	}
	// 判断余数为0的是否出现偶数次
	if mods[0]&1 == 1 {
		return false
	}
	for i, j := 1, k-1; i <= j; i, j = i+1, j-1 {
		// j = k - i
		// 判断余数为i的和余数为k-i的个数是否相同，可以组成一对
		if mods[i] != mods[j] {
			return false
		}
	}
	return true
}
