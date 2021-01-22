// Brute Force
//
// Your runtime beats 6.67 % of golang submissions (772 ms)
// Your memory usage beats 66.67 % of golang submissions (2.1 MB)

func countTriplets(arr []int) int {
	ret := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			xor1 := 0
			for ii := i; ii < j; ii++ {
				xor1 ^= arr[ii]
			}
			for k := j; k < len(arr); k++ {
				xor2 := 0
				for jj := j; jj <= k; jj++ {
					xor2 ^= arr[jj]
				}
				if xor1 == xor2 {
					ret += 1
				}
			}
		}
	}
	return ret
}
