// Your runtime beats 8.33 % of golang submissions (1793 ms)
// Your memory usage beats 66.67 % of golang submissions (9 MB)
//
// 因为子序列不要求连续，只关注子序列的最大值和最小值，所以不关心相对位置，可以将数组进行排序
// 使用双指针找出排序后连续的子序列最大值+最小值小于等于target的最大长度
// 左指针i，右指针j，满足条件的子序列为，最小值必须要有，其他元素都是可有可无的
// 所以子序列的个数为 2 ** (j - i - 1)
// 因为 2 ** (10 ^ 5) 会溢出，所以在计算幂的时候就需要进行取余
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)

var pow = make([]int, int(1e5))
var m int = 1e9 + 7

func numSubseq(nums []int, target int) int {
	var ret int = 0
	sort.Ints(nums)
	pow[0] = 1
	for i := 0; i < len(nums); i++ {
		if nums[i]+nums[i] > target {
			break
		}
		j := i
		for j < len(nums) && nums[i]+nums[j] <= target {
			j++
		}
		ret = (ret + getPow(j-i-1)) % m
	}
	return int(ret)
}

func getPow(i int) int {
	if pow[i] == 0 {
		pow[i] = 2 * getPow(i-1) % m
	}
	return pow[i]
}
