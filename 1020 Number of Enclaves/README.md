# [1020. 飞地的数量](https://leetcode-cn.com/problems/number-of-enclaves/)
## Description
给出一个二维数组 A，每个单元格为 0（代表海）或 1（代表陆地）。  
移动是指在陆地上从一个地方走到另一个地方（朝四个方向之一）或离开网格的边界。  
返回网格中**无法**在任意次数的移动中离开网格边界的陆地单元格的数量。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2021/02/18/enclaves1.jpg)
Input:  
```
[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
```
Output:
```
3
```
有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
### Example 2
![](https://assets.leetcode.com/uploads/2021/02/18/enclaves2.jpg)
Input:  
```
[[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
```
Output:
```
0
```
所有 1 都在边界上或可以到达边界。
## Hint
- 1 <= A.length <= 500
- 1 <= A[i].length <= 500
- 0 <= A[i][j] <= 1
- 所有行的大小都相同
