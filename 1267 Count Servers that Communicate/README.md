# [1267. 统计参与通信的服务器](https://leetcode-cn.com/problems/count-servers-that-communicate/)
## Description
这里有一幅服务器分布图，服务器的位置标识在 `m * n` 的整数矩阵网格 `grid` 中，1 表示单元格上有服务器，0 表示没有。  
如果两台服务器位于同一行或者同一列，我们就认为它们之间可以进行通信。  
请你统计并返回能够与至少一台其他服务器进行通信的服务器的数量。  
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-6.jpg)  
Input:  
```
grid = [[1,0],[0,1]]
```
Output:
```
0
```
没有一台服务器能与其他服务器进行通信。
### Example 2
![](https://assets.leetcode.com/uploads/2019/11/13/untitled-diagram-4.jpg)  
Input:  
```
grid = [[1,0],[1,1]]
```
Output:
```
3
```
所有这些服务器都至少可以与一台别的服务器进行通信。
### Example 3
![](https://assets.leetcode.com/uploads/2019/11/14/untitled-diagram-1-3.jpg)  
Input:  
```
grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
```
Output:
```
4
```
第一行的两台服务器互相通信，第三列的两台服务器互相通信，但右下角的服务器无法与其他服务器通信。
## Hint
- m == grid.length
- n == grid[i].length
- 1 <= m <= 250
- 1 <= n <= 250
- grid[i][j] == 0 or 1
