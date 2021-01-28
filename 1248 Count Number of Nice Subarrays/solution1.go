// Brute Force
// 
// Time Limit Exceeded
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)

func numberOfSubarrays(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i] & 1
		if sum == k {
			count += 1
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] & 1 == 1 {
				sum = sum + 1 
			} 
			if sum == k {
				count += 1
			}
			if sum > k {
				break
			}
		}
	}
	return count
}
