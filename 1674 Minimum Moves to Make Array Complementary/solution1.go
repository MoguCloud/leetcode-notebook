// Your runtime beats 100 % of golang submissions (120 ms)
// Your memory usage beats 50 % of golang submissions (10.7 MB)
//
// 差分数组
//
// 假设 a = nums[i] ,  b = nums[n-1-i]
// a + b 最小值为 2， 即都变成1时候
// a + b 只需要改动 1 个数时候，可以变成的最小值为 1 + min(a, b), 最大值为 max(a, b) + limit
// 那么 a + b 在 [2, 1 + min(a, b) ) 范围内，则需要改动 2 个数
// a + b 在 [1 + min(a, b), max(a, b) + limit] 范围内，只需要改动 1 个数，但是需要排除 = a + b 的情况，因为此时不需要改动
// a + b 在 (max(a, b) + limit, 2 * limit] 范围内，需崖改动 2 个数
//
//  2 <= sum < 1 + min(a, b)            操作数 = 2
//  1 + min(a, b) <= sum < a + b        操作数 = 1
//  sum = a + b                         操作数 = 0
//  a + b < sum <= max(a, b) + limit    操作数 = 1
//  max(a, b) + limit < sum <= 2*limit  操作数 = 2
//
//
// 差分数组
// arr	  2 4 5 7 9
// index  0 1 2 3 4
// diff   2 2 1 2 2
//
// arr    2 5 6 8 9
// diff   2 3 1 2 1
//
// 差分数组每一项 diff[i] = arr[i] - arr[i-1]
// arr[i:j] += delta => diff[i] += delta && diff[j+1] -= delta
// 用差分数组推回原数组，只要累加即可
//
// 时间复杂度 O(n)
// 空间复杂度 O(limit)

func minMoves(nums []int, limit int) int {
	diff := make([]int, 2*limit+2)
	for i := 0; i < len(nums)/2; i++ {
		a := nums[i]
		b := nums[len(nums)-i-1]
		// 2 <= sum < 1 + min(a, b)            操作数 = 2
		diff[2] += 2
		diff[1+min(a, b)] -= 2
		// 1 + min(a, b) <= sum < a + b        操作数 = 1
		diff[1+min(a, b)] += 1
		diff[a+b] -= 1
		// a + b < sum <= max(a, b) + limit    操作数 = 1
		diff[a+b+1] += 1
		diff[max(a, b)+limit+1] -= 1
		// max(a, b) + limit < sum <= 2*limit  操作数 = 2
		diff[max(a, b)+limit+1] += 2
		diff[2*limit+1] -= 2
	}
	op := 0
	ret := diff[2]
	for i := 2; i <= 2*limit; i++ {
		op += diff[i]
		ret = min(ret, op)
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
