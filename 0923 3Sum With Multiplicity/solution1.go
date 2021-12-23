// Your runtime beats 11.11 % of golang submissions (1300 ms)
// Your memory usage beats 11.11 % of golang submissions (7.1 MB)
//
// 哈希表
// i 从 [0, len(arr)-1]开始遍历，内层循环中的 j, k 的和2sum处理方式一样
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

func threeSumMulti(arr []int, target int) int {
	MOD := 1000000007
	ret := 0
	for i := 0; i < len(arr); i++ {
		memo := make(map[int]int)
		for j := i + 1; j < len(arr); j++ {
			if count, ok := memo[arr[j]]; ok {
				ret += count
				ret %= MOD
			}
			memo[target-arr[i]-arr[j]] += 1
		}
	}
	return ret
}
