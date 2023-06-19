# [2348. 全 0 子数组的数目](https://leetcode.cn/problems/number-of-zero-filled-subarrays/)
## Description
给你一个整数数组 `nums` ，返回全部为 `0` 的 **子数组** 数目。  
**子数组** 是一个数组中一段连续非空元素组成的序列。
## Example
### Example 1
Input:  
```
nums = [1,3,0,0,2,0,0,4]
```
Output:
```
6
```
子数组 [0] 出现了 4 次。  
子数组 [0,0] 出现了 2 次。  
不存在长度大于 2 的全 0 子数组，所以我们返回 6 。  
### Example 2
Input:  
```
nums = [0,0,0,2,0,0]
```
Output:
```
9
```
子数组 [0] 出现了 5 次。   
子数组 [0,0] 出现了 3 次。  
子数组 [0,0,0] 出现了 1 次。  
不存在长度大于 3 的全 0 子数组，所以我们返回 9 。
### Example 3
Input:  
```
nums = [2,10,2019]
```
Output:
```
0
```
没有全 0 子数组，所以我们返回 0 。
## Hint
- 1 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9

