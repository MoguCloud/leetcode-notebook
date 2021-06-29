# [525. 连续数组](https://leetcode-cn.com/problems/contiguous-array/)
## Description
给定一个二进制数组 `nums` , 找到含有相同数量的 `0` 和 `1` 的最长连续子数组，并返回该子数组的长度。
## Example
### Example 1
Input:  
```
nums = [0,1]
```
Output:
```
2
```
[0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
### Example 2
Input:  
```
nums = [0,1,0]
```
Output:
```
2
```
[0, 1]  (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
## Hint
- 1 <= nums.length <= 105
- nums[i] 不是 0 就是 1
