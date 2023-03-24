# [994. 腐烂的橘子](https://leetcode.cn/problems/rotting-oranges/)
## Description
在给定的 `m x n` 网格 `grid` 中，每个单元格可以有以下三个值之一：
- 值 `0` 代表空单元格；  
- 值 `1` 代表新鲜橘子；  
- 值 `2` 代表腐烂的橘子。  
每分钟，腐烂的橘子 **周围 4 个方向上相邻** 的新鲜橘子都会腐烂。  
返回 **直到单元格中没有新鲜橘子为止所必须经过的最小分钟数**。如果不可能，返回 `-1` 。  
## Example
### Example 1
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/02/16/oranges.png)  
Input:  
```
grid = [[2,1,1],[1,1,0],[0,1,1]]
```
Output:
```
4
```
### Example 2
Input:  
```
grid = [[2,1,1],[0,1,1],[1,0,1]]
```
Output:
```
-1
```
左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
### Example 3
Input:  
```
grid = [[0,2]]
```
Output:
```
0
```
因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
## Hint
- m == grid.length
- n == grid[i].length
- 1 <= m, n <= 10
- grid[i][j] 仅为 0、1 或 2

