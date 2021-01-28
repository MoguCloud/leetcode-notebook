# [1248. 统计「优美子数组」](https://leetcode-cn.com/problems/count-number-of-nice-subarrays/)
## Description
给你一个整数数组 `nums` 和一个整数 `k`。  
如果某个 **连续** 子数组中恰好有 `k` 个奇数数字，我们就认为这个子数组是「**优美子数组**」。  
请返回这个数组中「优美子数组」的数目。  
## Example
### Example 1
Input:  
```
nums = [1,1,2,1,1], k = 3
```
Output:
```
2
```
包含 3 个奇数的子数组是 [1,1,2,1] 和 [1,2,1,1] 。
### Example 2
Input:  
```
nums = [2,4,6], k = 1
```
Output:
```
0
```
数列中不包含任何奇数，所以不存在优美子数组。
### Example 3
Input:  
```
nums = [2,2,2,1,2,2,1,2,2,2], k = 2
```
Output:
```
16
```
## Hint
- 1 <= nums.length <= 50000
- 1 <= nums[i] <= 10^5
- 1 <= k <= nums.length
