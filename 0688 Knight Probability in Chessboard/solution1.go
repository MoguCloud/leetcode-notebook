// Force Brute
//
// Time Limit Exceeded
// O(8^k)

// 移动K次后还在棋盘的数量
var probableCount int

// 棋盘长度N
var board int

// 总移动次数K
var move int

type knightStruct struct {
	// x坐标
	x int
	// y坐标
	y int
	// 移动次数
	m int
}

var knight *knightStruct

func knightProbability(N int, K int, r int, c int) float64 {
	probableCount = 0
	board = N
	move = K
	knight = &knightStruct{
		x: r,
		y: c,
		m: 0,
	}
	check()
	ret := float64(probableCount) / math.Pow(8, float64(K))
	return ret
}

func check() {
	if knight.x < 0 || knight.x >= board || knight.y < 0 || knight.y >= board {
		return
	}
	if knight.m < move {
		knight.x += 2
		knight.y += 1
		knight.m += 1
		check()
		knight.x -= 2
		knight.y -= 1
		knight.m -= 1

		knight.x += 1
		knight.y += 2
		knight.m += 1
		check()
		knight.x -= 1
		knight.y -= 2
		knight.m -= 1

		knight.x -= 2
		knight.y -= 1
		knight.m += 1
		check()
		knight.x += 2
		knight.y += 1
		knight.m -= 1

		knight.x -= 1
		knight.y -= 2
		knight.m += 1
		check()
		knight.x += 1
		knight.y += 2
		knight.m -= 1

		knight.x += 2
		knight.y -= 1
		knight.m += 1
		check()
		knight.x -= 2
		knight.y += 1
		knight.m -= 1

		knight.x += 1
		knight.y -= 2
		knight.m += 1
		check()
		knight.x -= 1
		knight.y += 2
		knight.m -= 1

		knight.x -= 2
		knight.y += 1
		knight.m += 1
		check()
		knight.x += 2
		knight.y -= 1
		knight.m -= 1

		knight.x -= 1
		knight.y += 2
		knight.m += 1
		check()
		knight.x += 1
		knight.y -= 2
		knight.m -= 1

	} else if knight.m == move {
		probableCount += 1
	}
}
