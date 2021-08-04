// Your runtime beats 20 % of golang submissions (48 ms)
// Your memory usage beats 100 % of golang submissions (7.1 MB)
//
// 左边连续的1不进行移动
// 操作1 可以将左边连续的n个0 变成 n-1个1和1个0
// 操作2 可以将左边的1移到最后边
// 最后会变成 11..101...1的形态
// 左侧的1可能为0个 0最多1个
// 左侧1的个数为一开始左侧连续1的个数 + 0的个数 - 1

func maximumBinaryString(binary string) string {
	leftOneCount := 0
	leftFlag := binary[0] == '1'
	zeroCount := 0
	for i := 0; i < len(binary); i++ {
		if binary[i] == '0' {
			zeroCount += 1
			leftFlag = false
		}
		if binary[i] == '1' && leftFlag {
			leftOneCount += 1
		}
	}
	ret := make([]byte, len(binary))
	for i := 0; i < len(ret); i++ {
		ret[i] = '1'
	}
	if zeroCount > 0 {
		ret[leftOneCount+zeroCount-1] = '0'
	}
	return string(ret)
}
