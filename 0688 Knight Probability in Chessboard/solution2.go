// 带记忆的 DFS
// Your runtime beats 12.12 % of golang submissions (24 ms)
// Your memory usage beats 6.06 % of golang submissions (6.6 MB)

// 棋盘长度N
var board int

// 总移动次数K
var move int

type status struct {
	r int
	c int
	k int
}

var memo map[status]float64

func knightProbability(N int, K int, r int, c int) float64 {
	board = N
	move = K
	memo = make(map[status]float64)
	probableCount := dfs(r, c, 0)
	return probableCount / math.Pow(8, float64(K))
}

func dfs(r int, c int, k int) float64 {
	s := status{
		r: r,
		c: c,
		k: k,
	}
	if count, ok := memo[s]; ok {
		return count
	}
	if k == move {
		memo[s] = 1
		return 1
	} else {
		var count float64 = 0
		if r+2 < board {
			if c+1 < board {
				count += dfs(r+2, c+1, k+1)
			}
			if c-1 >= 0 {
				count += dfs(r+2, c-1, k+1)
			}
		}
		if r-2 >= 0 {
			if c+1 < board {
				count += dfs(r-2, c+1, k+1)
			}
			if c-1 >= 0 {
				count += dfs(r-2, c-1, k+1)
			}
		}
		if r+1 < board {
			if c+2 < board {
				count += dfs(r+1, c+2, k+1)
			}
			if c-2 >= 0 {
				count += dfs(r+1, c-2, k+1)
			}
		}
		if r-1 >= 0 {
			if c+2 < board {
				count += dfs(r-1, c+2, k+1)
			}
			if c-2 >= 0 {
				count += dfs(r-1, c-2, k+1)
			}
		}
		memo[s] = count
		return count
	}
	return 0
}
