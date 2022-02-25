// Time Limit Exceeded
//
// DFS

var memo map[int]struct{}

func canReach(s string, minJump int, maxJump int) bool {
	memo = make(map[int]struct{})
	return dfs(s, minJump, maxJump, 0)
}

func dfs(s string, minJump int, maxJump int, i int) bool {
	if i == len(s)-1 {
		return true
	}
	if _, ok := memo[i]; ok {
		return false
	}
	memo[i] = struct{}{}
	for j := min(i+maxJump, len(s)-1); j >= i+minJump; j-- {
		if s[j] == '0' {
			if dfs(s, minJump, maxJump, j) {
				return true
			}
		}
	}
	return false
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
