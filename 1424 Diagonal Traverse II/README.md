# [1424. 对角线遍历 II](https://leetcode-cn.com/problems/diagonal-traverse-ii/)
## Description
给你一个列表 `nums` ，里面每一个元素都是一个整数列表。请你依照下面各图的规则，按顺序返回 `nums` 中对角线上的整数。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2020/04/08/sample_1_1784.png)  
Input:  
```
nums = [[1,2,3],[4,5,6],[7,8,9]]
```
Output:
```
[1,4,2,7,5,3,8,6,9]
```
### Example 2
![](https://assets.leetcode.com/uploads/2020/04/08/sample_2_1784.png)  
Input:  
```
[[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
```
Output:
```
[1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
```
### Example 3
Input:  
```
[[1,2,3],[4],[5,6,7],[8],[9,10,11]]
```
Output:
```
[1,4,2,5,3,8,6,9,7,10,11]
```
### Example 4
Input:  
```
[[1,2,3,4,5,6]]
```
Output:
```
[[1,2,3,4,5,6]]
```
## Hint
- 1 <= nums.length <= 10^5
- 1 <= nums[i].length <= 10^5
- 1 <= nums[i][j] <= 10^9
- nums 中最多有 10^5 个数字。
