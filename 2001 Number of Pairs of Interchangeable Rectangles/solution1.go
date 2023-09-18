// Your runtime beats 53.33 % of golang submissions (308 ms)
// Your memory usage beats 20.00 % of golang submissions (26.06 MB)
//
// Hash Table
// 使用hash table记录相同长宽比的数量
// 直接使用长除以宽有精度问题
// 所以使用长和宽除以最大公约数作为key

func interchangeableRectangles(rectangles [][]int) int64 {
	counts := make(map[string]int)
	for _, r := range rectangles {
		k := gcd(r[0], r[1])
		counts[fmt.Sprintf("%d %d", r[0]/k, r[1]/k)]++
	}
	var ret int64
	for _, count := range counts {
		ret += int64(count) * int64(count-1) / 2
	}
	return ret
}

func gcd(i int, j int) int {
	if i < j {
		i, j = j, i
	}
	mod := i % j
	if mod == 0 {
		return j
	} else {
		return gcd(j, mod)
	}

}
