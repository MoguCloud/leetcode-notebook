// Your runtime beats 94.12 % of golang submissions (52 ms)
// Your memory usage beats 100 % of golang submissions (6.5 MB)
//
// 分别统计每一行，每一列服务器的总数。对于一台服务器，它所在的行或者列上的服务器数大于1，则该服务器有可以通信的其他服务器。

func countServers(grid [][]int) int {
	rowCount := make([]int, len(grid))
	columnCount := make([]int, len(grid[0]))
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				rowCount[i] += 1
				columnCount[j] += 1
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && (rowCount[i] > 1 || columnCount[j] > 1) {
				count += 1
			}
		}
	}
	return count
}
