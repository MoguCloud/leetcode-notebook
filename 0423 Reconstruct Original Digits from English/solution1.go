// Your runtime beats 50 % of golang submissions (12 ms)
// Your memory usage beats 25 % of golang submissions (5.3 MB)
//
// 记录每个字母出现的次数，根据字母特征来判断数字出现的次数
// zero    e        o r        z
// one     e      n o
// two              o   t     w
// three  2e   h      r t 
// four      f      o r  u
// five    e f  i         v
// six          i      s     x 
// seven  2e      n    s  v 
// eight   e  ghi       t
// nine    e    i2n
// => num(zero) num(z) 
//    num(two) = num(w) 
//    num(six) = num(x)
//    num(eight) = num(g)
//    num(four) = num(u)
//    num(five) = num(f) - num(four)
//    num(seven) = num(v) - num(five)
//    num(three) = num(h) - num(eight)
//    num(nine) = num(i) - num(six) - num(eight) - num(five)
//    num(one) = num(o) - num(zero) - num(two) - num(four)
func originalDigits(s string) string {
	numArr := make([]int, 10)
	numMap := make(map[byte]int)
	var n int
	for i := 0; i < len(s); i++ {
		numMap[s[i]] += 1
	}
	n, _ = numMap['z']
	numArr[0] = n
	n, _ = numMap['w']
	numArr[2] = n
	n, _ = numMap['x']
	numArr[6] = n
	n, _ = numMap['g']
	numArr[8] = n
	n, _ = numMap['u']
	numArr[4] = n
	n, _ = numMap['f']
	numArr[5] = n - numArr[4]
	n, _ = numMap['v']
	numArr[7] = n - numArr[5]
	n, _ = numMap['h']
	numArr[3] = n - numArr[8]
	n, _ = numMap['i']
	numArr[9] = n - numArr[6] - numArr[8] - numArr[5]
	n, _ = numMap['o']
	numArr[1] = n - numArr[0] - numArr[2] - numArr[4]
	ret := []byte{}
	for i := 0; i < len(numArr); i++ {
		for j := 0; j < numArr[i]; j++ {
			ret = append(ret, byte('0' + i))
		}
	}
	return string(ret)
}
