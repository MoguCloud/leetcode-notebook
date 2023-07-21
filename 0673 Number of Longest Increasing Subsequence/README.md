# [673. 最长递增子序列的个数](https://leetcode.cn/problems/number-of-longest-increasing-subsequence/)
## Description
给定一个未排序的整数数组 `nums` ， *返回最长递增子序列的个数* 。  
**注意** 这个数列必须是 **严格** 递增的。
## Example
### Example 1
Input:  
```
[1,3,5,4,7]
```
Output:
```
2
```
有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
### Example 2
Input:  
```
[2,2,2,2,2]
```
Output:
```
5
```
最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
## Hint
- 1 <= nums.length <= 2000
- -10^6 <= nums[i] <= 10^6
