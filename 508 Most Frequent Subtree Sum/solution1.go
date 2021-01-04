var maxCount int

// key sum
// value 出现次数
var countMap map[int]int

func findFrequentTreeSum(root *TreeNode) []int {
	countMap = make(map[int]int)
	maxCount = 0

	ret := []int{}

	if root != nil {
		calcSum(root)
	}
	for k, v := range countMap {
		if v == maxCount {
			ret = append(ret, k)
		}
	}
	return ret

}

func calcSum(root *TreeNode) int {
	left := 0
	right := 0
	if root.Left != nil {
		left = calcSum(root.Left)
	}
	if root.Right != nil {
		right = calcSum(root.Right)
	}
	sum := root.Val + left + right
	checkSum(sum)
	return sum
}

func checkSum(sum int) {
	count, _ := countMap[sum]
	count += 1
	countMap[sum] = count
	if count > maxCount {
		maxCount = count
	}
}
