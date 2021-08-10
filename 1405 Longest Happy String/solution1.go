// Time Limit Exceeded
//
// BackTracking

var ret string
var byteArr []byte
var stop bool

func longestDiverseString(a int, b int, c int) string {
	stop = false
	ret = ""
	byteArr = []byte{}
	backtrack(a, b, c, 0, 0, 0)
	return ret
}

func backtrack(a int, b int, c int, aa int, bb int, cc int) {
	if stop {
		return
	}
	if len(byteArr) > len(ret) {
		ret = string(byteArr)
	}
	if aa == a && bb == b && cc == c {
		stop = true
	} else {
		if aa < a && !(len(byteArr) >= 2 && byteArr[len(byteArr)-1] == 'a' && byteArr[len(byteArr)-2] == 'a') {
			byteArr = append(byteArr, 'a')
			backtrack(a, b, c, aa+1, bb, cc)
			byteArr = byteArr[:len(byteArr)-1]
		}
		if bb < b && !(len(byteArr) >= 2 && byteArr[len(byteArr)-1] == 'b' && byteArr[len(byteArr)-2] == 'b') {
			byteArr = append(byteArr, 'b')
			backtrack(a, b, c, aa, bb+1, cc)
			byteArr = byteArr[:len(byteArr)-1]
		}
		if cc < c && !(len(byteArr) >= 2 && byteArr[len(byteArr)-1] == 'c' && byteArr[len(byteArr)-2] == 'c') {
			byteArr = append(byteArr, 'c')
			backtrack(a, b, c, aa, bb, cc+1)
			byteArr = byteArr[:len(byteArr)-1]
		}
	}
}
