# [377. 组合总和 Ⅳ](https://leetcode.cn/problems/combination-sum-iv/description/)
## Description
给你一个由 **不同** 整数组成的数组 `nums` ，和一个目标整数 `target` 。请你从 `nums` 中找出并返回总和为 `target` 的元素组合的个数。  
题目数据保证答案符合 32 位整数范围。
## Example
### Example 1
Input:  
```
nums = [1,2,3], target = 4
```
Output:
```
7
```
所有可能的组合为：  
(1, 1, 1, 1)  
(1, 1, 2)  
(1, 2, 1)  
(1, 3)  
(2, 1, 1)  
(2, 2)  
(3, 1)  
请注意，顺序不同的序列被视作不同的组合。
### Example 2
Input:  
```
nums = [9], target = 3
```
Output:
```
0
```
## Hint
- 1 <= nums.length <= 200
- 1 <= nums[i] <= 1000
- nums 中的所有元素 互不相同
- 1 <= target <= 1000
