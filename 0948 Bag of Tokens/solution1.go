// 贪心
//
// Your runtime beats 100 % of golang submissions (4 ms)
// Your memory usage beats 100 % of golang submissions (3 MB)
//
// P点不够用了就用1个score换成最大面值的token
// 然后用P点换尽可能多的score，从最小面值的token开始，直到换不起最小剩余面值的token

func bagOfTokensScore(tokens []int, P int) int {
	maxScore := 0
	score := 0
	sort.Ints(tokens)
	i := 0
	j := len(tokens) - 1
	for i <= j {
		if P >= tokens[i] {
			P = P - tokens[i]
			i += 1
			score += 1
			if score > maxScore {
				maxScore = score
			}
			continue
		}
		if score > 0 {
			P = P + tokens[j]
			score -= 1
			j -= 1
		} else {
			break
		}
	}
	return maxScore
}
