# [1218. 最长定差子序列](https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/)
## Description
给你一个整数数组 `arr` 和一个整数 `difference`，请你找出并返回 `arr` 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 `difference` 。  
**子序列** 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 `arr` 派生出来的序列。
## Example
### Example 1
Input:  
```
arr = [1,2,3,4], difference = 1
```
Output:
```
4
```
最长的等差子序列是 [1,2,3,4]。
### Example 2
Input:  
```
arr = [1,3,5,7], difference = 1
```
Output:
```
1
```
最长的等差子序列是任意单个元素。
### Example 3
Input:  
```
arr = [1,5,7,8,5,3,4,2,1], difference = -2
```
Output:
```
4
```
最长的等差子序列是 [7,5,3,1]。
## Hint
- 1 <= arr.length <= 105
- -104 <= arr[i], difference <= 104

