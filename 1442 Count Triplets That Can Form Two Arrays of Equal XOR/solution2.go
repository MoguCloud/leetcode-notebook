// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 66.67 % of golang submissions (2.1 MB)
//
// a == b 等同于 a ^ b = 0
// a = arr[i] ^ ... ^ arr[j - 1]
// b = arr[j] ^ ... ^ arr[k]
// 所以 a == b 等同于 arr[i] ^ ... arr[k] = 0 
// 此时 j 只要满足 i < j <= k 都能使得 a == b
//
// 时间复杂度 O(n ^ 2)
// 空间复杂度 O(1)

func countTriplets(arr []int) int {
	ret := 0
	for i := 0; i < len(arr)-1; i++ {
		n := arr[i]
		for j := i + 1; j < len(arr); j++ {
			n ^= arr[j]
			if n == 0 {
				ret += j - i
			}
		}
	}
	return ret
}
