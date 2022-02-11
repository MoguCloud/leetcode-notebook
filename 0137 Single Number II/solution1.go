// Your runtime beats 45.31 % of golang submissions (8 ms)
// Your memory usage beats 84.38 % of golang submissions (3.4 MB)
//
// nums [3, 3, 3, 5]
// 二进制 3    0 1 1
// 二进制 3    0 1 1
// 二进制 3    0 1 1
// 二进制 5    1 0 1
// 按位求和    1 3 4
// 和 % 3     1 0 1  = 5

func singleNumber(nums []int) int {
	ret := int32(0)
	for i := 0; i < 32; i++ {
		sum := int32(0)
		for _, num := range nums {
			sum += (int32(num) >> i) & 1
		}
		ret |= (sum % 3) << i
	}
	return int(ret)
}
