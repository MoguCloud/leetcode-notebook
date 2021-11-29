// Your runtime beats 89.58 % of golang submissions (8 ms)
// Your memory usage beats 79.17 % of golang submissions (6.5 MB)
//
// Stack
//
// 用栈记录函数调用，遇到start压入栈中，遇到end从栈中弹出
// start 是从「这1秒的开始」开始计算
// end 是从「这一秒的结束」开始计算
// 所以遇到end时候进行时间计算需要+1
// 遇到end是需要将栈顶元素的时间改为当前时间的下一秒，即将上一个调用进行继续运行，
// 因为从当前时间的下一秒开始计数，而不是从压入栈的时间进行计算。

type call struct {
	funcID    int
	timestamp int
}

func exclusiveTime(n int, logs []string) []int {
	ret := make([]int, n)
	stack := []call{}
	for _, log := range logs {
		s := strings.Split(log, ":")
		funcID, _ := strconv.Atoi(s[0])
		status := s[1]
		timestamp, _ := strconv.Atoi(s[2])
		if status == "start" {
			if len(stack) > 0 {
				runTime := timestamp - stack[len(stack)-1].timestamp
				ret[stack[len(stack)-1].funcID] += runTime
			}
			stack = append(stack, call{funcID: funcID, timestamp: timestamp})
		} else {
			runTime := timestamp - stack[len(stack)-1].timestamp + 1
			ret[stack[len(stack)-1].funcID] += runTime
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				stack[len(stack)-1].timestamp = timestamp + 1
			}
		}
	}
	return ret
}
