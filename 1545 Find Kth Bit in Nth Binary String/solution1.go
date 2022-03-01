// Your runtime beats 50 % of golang submissions (3 ms)
// Your memory usage beats 92.86 % of golang submissions (2 MB)

func findKthBit(n int, k int) byte {
	length := 1
	for length <= k {
		length = length*2 + 1
	}
	rev := 0
	for k > 1 {
		if k <= (length+1)/2 {
			length = length / 2
		} else {
			rev += 1
			k = length - k + 1
			length = length / 2
		}
	}
	if rev&1 == 1 {
		return '1'
	} else {
		return '0'
	}
}
