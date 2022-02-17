# [454. 四数相加 II](https://leetcode-cn.com/problems/4sum-ii/)
## Description
给你四个整数数组 `nums1`、`nums2`、`nums3` 和 `nums4` ，数组长度都是 `n` ，请你计算有多少个元组 `(i, j, k, l)` 能满足：
- 0 <= i, j, k, l < n
- nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
## Example
### Example 1
Input:  
```
nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
```
Output:
```
2
```
两个元组如下：  
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0  
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0  
### Example 2
Input:  
```
nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
```
Output:
```
1
```
## Hint
- n == nums1.length
- n == nums2.length
- n == nums3.length
- n == nums4.length
- 1 <= n <= 200
- -2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28