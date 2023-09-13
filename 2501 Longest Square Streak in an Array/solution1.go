// Your runtime beats 80.00 % of golang submissions (188 ms)
// Your memory usage beats 100 % of golang submissions (10.12 MB)
//
// Hash Table
// 将数组中元素存入哈希表中
// 将数组元素进行排序
// 依次查找该元素的平方是否存在，存在的话则继续寻找平方
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

func longestSquareStreak(nums []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = false
	}
	sort.Ints(nums)
	ret := 0
	for _, num := range nums {
		count := 1
		sqr := num
		for {
			sqr = sqr * sqr
			visited, ok := numMap[sqr]
			if visited {
				break
			}
			if ok {
				count += 1
				numMap[sqr] = true
			} else {
				break
			}
		}
		if count > ret {
			ret = count
		}
	}
	if ret == 1 {
		ret = -1
	}
	return ret
}
