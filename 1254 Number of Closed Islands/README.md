# [1254. 统计封闭岛屿的数目](https://leetcode.cn/problems/number-of-closed-islands/)
## Description
二维矩阵 `grid` 由 `0` （土地）和 `1` （水）组成。岛是由最大的4个方向连通的 `0` 组成的群，封闭岛是一个 `完全` 由1包围（左、上、右、下）的岛。  
请返回 **封闭岛屿** 的数目。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2019/10/31/sample_3_1610.png)   
Input:  
```
grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
```
Output:
```
2
```
灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
### Example 2
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/11/07/sample_4_1610.png)  
Input:  
```
grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
```
Output:
```
1
```
### Example 3
Input:
```
grid = [[1,1,1,1,1,1,1],
        [1,0,0,0,0,0,1],
        [1,0,1,1,1,0,1],
        [1,0,1,0,1,0,1],
        [1,0,1,1,1,0,1],
        [1,0,0,0,0,0,1],
        [1,1,1,1,1,1,1]]
```
Output:
```
2
```
## Hint:
- 1 <= grid.length, grid[0].length <= 100
- 0 <= grid[i][j] <=1
