// Your runtime beats 44.83 % of golang submissions (36 ms)
// Your memory usage beats 62.07 % of golang submissions (7.3 MB)
//
// DFS
//
// 时间复杂度 O(mn)
// 空间复杂度 O(mn)

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	dfs(&board, click)
	return board
}

func dfs(board *[][]byte, click []int) {
	if !checkRange(board, click[0], click[1]) {
		return
	}
	if (*board)[click[0]][click[1]] != 'E' {
		return
	}
	mineCount := 0
	for offsetX := -1; offsetX <= 1; offsetX++ {
		x := click[0] + offsetX
		for offsetY := -1; offsetY <= 1; offsetY++ {
			y := click[1] + offsetY
			if checkRange(board, x, y) && (*board)[x][y] == 'M' {
				mineCount += 1
			}
		}
	}
	if mineCount > 0 {
		(*board)[click[0]][click[1]] = byte('0' + mineCount)
	} else {
		(*board)[click[0]][click[1]] = 'B'
		for offsetX := -1; offsetX <= 1; offsetX++ {
			x := click[0] + offsetX
			for offsetY := -1; offsetY <= 1; offsetY++ {
				y := click[1] + offsetY
				dfs(board, []int{x, y})
			}
		}
	}
	return
}

func checkRange(board *[][]byte, x int, y int) bool {
	if x < 0 || x >= len(*board) || y < 0 || y >= len((*board)[0]) {
		return false
	} else {
		return true
	}
}
