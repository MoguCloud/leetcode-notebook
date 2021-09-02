# [1567. 乘积为正数的最长子数组长度](https://leetcode-cn.com/problems/maximum-length-of-subarray-with-positive-product/)
## Description
给你一个整数数组 `nums` ，请你求出乘积为正数的最长子数组的长度。  
一个数组的子数组是由原数组中零个或者更多个连续数字组成的数组。  
请你返回乘积为正数的最长子数组长度。  
## Example
### Example 1
Input:  
```
nums = [1,-2,-3,4]
```
Output:
```
4
```
数组本身乘积就是正数，值为 24 。
### Example 2
Input:  
```
nums = [0,1,-2,-3,-4]
```
Output:
```
3
```
最长乘积为正数的子数组为 [1,-2,-3] ，乘积为 6 。  
注意，我们不能把 0 也包括到子数组中，因为这样乘积为 0 ，不是正数。
### Example 3
Input:  
```
nums = [-1,-2,-3,0,1]
```
Output:
```
2
```
乘积为正数的最长子数组是 [-1,-2] 或者 [-2,-3] 。
### Example 4
Input:  
```
nums = [-1,2]
```
Output:
```
1
```
### Example 5
Input:  
```
nums = [1,2,3,5,-6,4,0,10]
```
Output:
```
4
```
## Hint
- 1 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9

