# [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/description/)
## Description
给你一个整数数组 `nums` ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。  
**子数组** 是数组中的一个连续部分。  
## Example
### Example 1
Input:  
```
nums = [-2,1,-3,4,-1,2,1,-5,4]
```
Output:
```
6
```
连续子数组 [4,-1,2,1] 的和最大，为 6 。
### Example 2
Input:  
```
nums = [1]
```
Output:
```
1
```
### Example 3
Input:  
```
nums = [5,4,-1,7,8]
```
Output:
```
23
```
## Hint
- 1 <= nums.length <= 10^5
- -10^4 <= nums[i] <= 10^4
