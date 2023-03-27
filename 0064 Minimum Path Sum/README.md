# [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)
## Description
给定一个包含非负整数的 `m x n` 网格 `grid` ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。  
说明：每次只能向下或者向右移动一步。  
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2020/11/05/minpath.jpg)  
Input:  
```
grid = [[1,3,1],[1,5,1],[4,2,1]]
```
Output:
```
7
```
因为路径 1→3→1→1→1 的总和最小。
### Example 2
Input:  
```
grid = [[1,2,3],[4,5,6]]
```
Output:
```
12
```
## Hint
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 200
- 0 <= grid[i][j] <= 100

