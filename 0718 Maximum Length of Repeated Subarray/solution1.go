// Your runtime beats 12.66 % of golang submissions (1636 ms)
// Your memory usage beats 94.94 % of golang submissions (3.5 MB)
//
// Brute Force
//
// 时间复杂度 O(n^3)
// 空间复杂度 O(1)

func findLength(nums1 []int, nums2 []int) int {
	maxLength := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			length := 0
			for k := 0; ; k++ {
				if i+k >= len(nums1) || j+k >= len(nums2) {
					break
				}
				if nums1[i+k] != nums2[j+k] {
					break
				} else {
					length += 1
				}
			}
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}
