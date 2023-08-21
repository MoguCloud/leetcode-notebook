// Your runtime beats 85.71 % of golang submissions (22 ms)
// Your memory usage beats 85.71 % of golang submissions (6.82 MB)
//
// 双指针
// 因为 L 可以向左移动， R 可以向右移动，但是 L 和 R 不能互相越过
// 所以可以判断 start 和 target 中 L 和 R 的顺序是否一致
// 还因为 L 只能向左移动， R 只能向右移动
// 所以 start 中的 第 i 个 R 的下标 必须 小于等于 target 中 第 i 个 R 的下标
// 并且 start 中的 第 i 个 L 的下标 必须 大于等于 target 中 第 i 个 L 的下标
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func canChange(start string, target string) bool {
	i := 0
	j := 0
	for i < len(start) && j < len(target) {
		// 跳过 _ 直到出现 L 或者 R
		for i < len(start) && start[i] == '_' {
			i++
		}
		// 跳过 _ 直到出现 L 或者 R
		for j < len(target) && target[j] == '_' {
			j++
		}
		// 避免下标越界
		if i >= len(start) || j >= len(target) {
			break
		}
		// 如果相对顺序不一样则不满足条件
		if start[i] != target[j] {
			return false
		}
		// start 中的 R 的下标 必须 小于等于 target 中 R 的下标
		if start[i] == 'R' && target[j] == 'R' && i > j {
			return false
		}
		// start 中的 L 的下标 必须 大于等于 target 中 L 的下标
		if start[i] == 'L' && target[j] == 'L' && i < j {
			return false
		}
		i++
		j++
	}
	// 剩余项不能包含非 _ 项
	for i < len(start) {
		if start[i] != '_' {
			return false
		}
		i++
	}
	for j < len(target) {
		if target[j] != '_' {
			return false
		}
		j++
	}
	return true
}
