// Your runtime beats 63.64 % of golang submissions (152 ms)
// Your memory usage beats 63.64 % of golang submissions (10.5 MB)
//
// 用哈希表记录已经访问过的元素
// 每轮遍历嵌套时，如果当前元素出现在哈希表中，则该轮遍历结束，判断该轮嵌套数是否最大
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func arrayNesting(nums []int) int {
	visited := make(map[int]struct{}) // 记录已经出现过的元素
	ret := 0
	for _, num := range nums {
		if len(nums)-len(visited) <= ret {
			break
		}
		if _, ok := visited[num]; ok {
			continue
		}
		count := 0
		for {
			if _, ok := visited[num]; ok {
				break
			}
			visited[num] = struct{}{}
			num = nums[num]
			count += 1
		}
		if count > ret {
			ret = count
		}
	}
	return ret
}
