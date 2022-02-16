// Your runtime beats 80.33 % of golang submissions (120 ms)
// Your memory usage beats 53.67 % of golang submissions (6.8 MB)
//
// 将 nums1, nums2 分为1组； num3, nums4 分为1组
// 分别用2个hash表记录 nums1+nums2所有相加出现的结果次数 和 num3+nums4和的结果次数
// 再遍历hash1的key，在hash2中寻找-key，如果存在即说明存在和为0的结果，个数为hash1[key]*hash2[-key]
//
// 时间复杂度 O(n^2)
// 空间复杂度 O(n^2

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	memo1 := make(map[int]int)
	memo2 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1); j++ {
			var sum int
			sum = nums1[i] + nums2[j]
			memo1[sum] += 1
			sum = nums3[i] + nums4[j]
			memo2[-sum] += 1
		}
	}
	ret := 0
	for k, v1 := range memo1 {
		if v2, ok := memo2[k]; ok {
			ret += v1 * v2
		}
	}
	return ret
}
