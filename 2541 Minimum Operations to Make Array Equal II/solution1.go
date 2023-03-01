// Your runtime beats 96.84 % of golang submissions (106 ms)
// Your memory usage beats 63.16 % of golang submissions (10 MB)
//
// 当k=0时，数组无法修改，需要nums1和nums2相等时才能满足条件
// 当k!=0时，记数组元素差diff[i]=nums1[i]-nums2[i]，
//          如果abs(diff[i])不能被k整除，则无法将nums1[i]加减k来变成nums2[i]
//          如果abs(diff[i])小于k时，则也无法将nums[i]加减k来变成nums2[i]
//          如果sum(diff)不等于0，则也无法将nums1变成nums2。因为经过一加一减k的操作是不会改变sum(diff)的
//          当满足以上条件时，则通过sum(abs(diff))/k/2次操作将nums1改成nums2。因为经过一次操作后，sum(abs(diff))减了2次k。
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	if k == 0 {
		for i := 0; i < len(nums1); i++ {
			if nums1[i] != nums2[i] {
				return -1
			}
		}
		return 0
	}
	var diffSum int64 = 0
	var diffAbsSum int64 = 0
	for i := 0; i < len(nums1); i++ {
		diff := nums1[i] - nums2[i]
		if diff != 0 && (abs(diff) < k || (abs(diff)%k != 0)) {
			return -1
		}
		diffSum += int64(diff)
		diffAbsSum += int64(abs(diff))
	}
	if diffSum != 0 {
		return -1
	}
	return diffAbsSum / int64(k) / 2
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
