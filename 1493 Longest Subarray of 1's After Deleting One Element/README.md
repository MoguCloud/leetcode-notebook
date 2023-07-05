# [1493. 删掉一个元素以后全为 1 的最长子数组](https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/)
## Description
给你一个二进制数组 `nums` ，你需要从中删掉一个元素。  
请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。  
如果不存在这样的子数组，请返回 0 。
## Example
### Example 1
Input:  
```
nums = [1,1,0,1]
```
Output:
```
3
```
删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
### Example 2
Input:  
```
nums = [0,1,1,1,0,1,1,0,1
```
Output:
```
5
```
删掉位置 4 的数字后，[0,1,1,1,1,1,0,1] 的最长全 1 子数组为 [1,1,1,1,1] 。
### Example 3
Input:  
```
nums = [1,1,1]
```
Output:
```
2
```
你必须要删除一个元素。
## Hint
- 1 <= nums.length <= 10^5
- nums[i] 要么是 0 要么是 1 。

