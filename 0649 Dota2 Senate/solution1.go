// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 6.6 % of golang submissions (7.7 MB)
//
// 贪心
// 每次都禁止投票顺序下一位的敌方阵营，投票顺序后面没有敌方阵营的则选择从开头顺序选择地方阵营
// 直到没有敌方阵营则结束
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func predictPartyVictory(senate string) string {
	qR := []int{}
	qD := []int{}
	// 记录各阵营的投票顺序
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			qR = append(qR, i)
		} else {
			qD = append(qD, i)
		}
	}
	for len(qR) > 0 && len(qD) > 0 {
		// 投票顺序小的优先投票
		// 因为优先选择投票顺序下一位的敌方阵营，当后面没有敌方阵营时才从开头选择
		// 所以投票完可以将该位参议员放到队列最后，只要保证相对位置即可
		if qR[0] < qD[0] {
			qR = append(qR, qR[0]+len(senate))
		} else {
			qD = append(qD, qD[0]+len(senate))
		}
		qR = qR[1:]
		qD = qD[1:]
	}
	// 当有一方阵营为空时则结束
	if len(qR) == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}
