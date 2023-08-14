# [215. 数组中的第K个最大元素](https://leetcode.cn/problems/kth-largest-element-in-an-array/description/)
## Description
给定整数数组 `nums` 和整数 `k`，请返回数组中第 `k` 个最大的元素。  
请注意，你需要找的是数组排序后的第 `k` 个最大的元素，而不是第 `k` 个不同的元素。  
你必须设计并实现时间复杂度为 `O(n)` 的算法解决此问题。  
## Example
### Example 1
Input:  
```
[3,2,1,5,6,4], k = 2
```
Output:
```
5
```
### Example 2
Input:  
```
[3,2,3,1,2,4,5,5,6], k = 4
```
Output:
```
4
```
## Hint
- 1 <= k <= nums.length <= 10^5
- -10^4 <= nums[i] <= 10^4
