// Your runtime beats 100 % of golang submissions (20 ms)
// Your memory usage beats 100 % of golang submissions (3.4 MB)
//
// 滑动窗口
//
// 时间复杂度 O(n ^ 2)
// 空间复杂度 O(1)

// nums1 = [1, 6, 1, 2, 4]
// nums2 = [7, 1, 2, 9]
// 将 nums2 不断移动/对齐 然后调用 findMaxLength 求最长重复子数组
// 过程如下
//  1.          [1, 6, 1, 2, 4]
//              [7, 1, 2, 9]
//               ^ 基准位置 最长重复子数组 0
//  2.          [1, 6, 1, 2, 4]
//           [7, 1, 2, 9]
//               ^ 基准位置 最长重复子数组 1
//  3.          [1, 6, 1, 2, 4]
//        [7, 1, 2, 9]
//               ^ 基准位置 最长重复子数组 0
//  4.          [1, 6, 1, 2, 4]
//     [7, 1, 2, 9]
//               ^ 基准位置 最长重复子数组 0
//  5. 交换 nums1 和 nums2 再次从第1步开始
//  6.	        [7, 1, 2, 9]
//              [1, 6, 1, 2, 4]
//               ^ 基准位置 最长重复子数组 0
//  7.	        [7, 1, 2, 9]
//           [1, 6, 1, 2, 4]
//               ^ 基准位置 最长重复子数组 2
//  8.	        [7, 1, 2, 9]
//        [1, 6, 1, 2, 4]
//               ^ 基准位置 最长重复子数组 0
//  9.	        [7, 1, 2, 9]
//     [1, 6, 1, 2, 4]
//               ^ 基准位置 最长重复子数组 0
//  10.	        [7, 1, 2, 9]
//  [1, 6, 1, 2, 4]
//               ^ 基准位置 最长重复子数组 0

func findLength(nums1 []int, nums2 []int) int {
	ret := 0
	for i := 0; i < len(nums1); i++ {
		l := findMaxLength(&nums1, &nums2, i, 0)
		if l > ret {
			ret = l
		}
	}
	for i := 0; i < len(nums2); i++ {
		l := findMaxLength(&nums1, &nums2, 0, i)
		if l > ret {
			ret = l
		}
	}
	return ret
}

// 计算 nums1[offset1:] 与 nums2[offset2:] 对齐后的最长重复子数组
func findMaxLength(nums1 *[]int, nums2 *[]int, offset1 int, offset2 int) int {
	minL := 0
	if len(*nums1)-offset1 < len(*nums2)-offset2 {
		minL = len(*nums1) - offset1
	} else {
		minL = len(*nums2) - offset2
	}
	maxL := 0
	length := 0
	for i := 0; i < minL; i++ {
		if (*nums1)[offset1+i] == (*nums2)[offset2+i] {
			length += 1
		} else {
			length = 0
		}
		if length > maxL {
			maxL = length
		}
	}
	return maxL
}
