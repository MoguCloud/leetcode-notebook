// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 39.02 % of golang submissions (2.8 MB)
//
// 模拟每天的变化，直到出现循环，根据天数和循环周期，取出对应的天数即可。

func prisonAfterNDays(cells []int, n int) []int {
	if n == 0 {
		return cells
	}
	cellArr := []string{}
	var newCells []int
	for i := 0; i < n; i++ {
		newCells = []int{}
		newCells = append(newCells, 0)
		for j := 1; j < 7; j++ {
			newCells = append(newCells, 1^cells[j-1]^cells[j+1])
		}
		newCells = append(newCells, 0)
		cells = newCells
		if len(cellArr) > 0 && sliceToStr(newCells) == cellArr[0] {
			break
		}
		cellArr = append(cellArr, sliceToStr(newCells))
	}
	remainder := n % len(cellArr)
	if remainder == 0 {
		return strToSlice(cellArr[len(cellArr)-1])
	}
	return strToSlice(cellArr[remainder-1])
}

func sliceToStr(nums []int) string {
	bytes := []byte{}
	for i := 0; i < len(nums); i++ {
		bytes = append(bytes, byte(nums[i]))
	}
	return string(bytes)
}

func strToSlice(strs string) []int {
	ret := []int{}
	for i := 0; i < len(strs); i++ {
		if strs[i] == 0 {
			ret = append(ret, 0)
		} else {
			ret = append(ret, 1)
		}
	}
	return ret
}
