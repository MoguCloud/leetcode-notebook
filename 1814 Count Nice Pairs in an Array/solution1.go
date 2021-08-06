// Time Limit Exceeded
// 76/84 cases passed (N/A)
//
// Brute Force

func countNicePairs(nums []int) int {
	var ret64 int64
	revArr := []int{}
	for i := 0; i < len(nums); i++ {
		rev := 0
		n := nums[i]
		for n > 0 {
			rev = rev*10 + n%10
			n = n / 10
		}
		revArr = append(revArr, rev)
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+revArr[i] == revArr[j]+nums[i] {
				ret64 += 1
			}
		}
	}
	return int(ret64 % 1000000007)
}
