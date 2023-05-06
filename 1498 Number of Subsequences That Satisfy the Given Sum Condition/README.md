# [1498. 满足条件的子序列数目](https://leetcode.cn/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/)  
## Description
给你一个整数数组 `nums` 和一个整数 `target` 。  
请你统计并返回 `nums` 中能满足其最小元素与最大元素的 **和** 小于或等于 `target` 的 **非空** 子序列的数目。  
由于答案可能很大，请将结果对 `10^9 + 7` 取余后返回。  
## Example
### Example 1
Input:  
```
nums = [3,5,6,7], target = 9
```
Output:
```
4
```
有 4 个子序列满足该条件。  
[3] -> 最小元素 + 最大元素 <= target (3 + 3 <= 9)  
[3,5] -> (3 + 5 <= 9)  
[3,5,6] -> (3 + 6 <= 9)  
[3,6] -> (3 + 6 <= 9)  
### Example 2
Input:  
```
nums = [3,3,6,8], target = 10
```
Output:
```
6
```
有 6 个子序列满足该条件。（nums 中可以有重复数字）  
[3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6]
### Example 3
Input:  
```
nums = [2,3,3,4,6,7], target = 12
```
Output:
```
61
```
共有 63 个非空子序列，其中 2 个不满足条件（[6,7], [7]）  
有效序列总数为（63 - 2 = 61）
## Hint
- 1 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^6
- 1 <= target <= 10^6

