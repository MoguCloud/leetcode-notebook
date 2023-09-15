# [1584. 连接所有点的最小费用](https://leetcode.cn/problems/min-cost-to-connect-all-points/description/)
## Description
给你一个`points` 数组，表示 2D 平面上的一些点，其中 `points[i] = [xi, yi]` 。  
连接点 `[xi, yi]` 和点 `[xj, yj]` 的费用为它们之间的 **曼哈顿距离** ：`|xi - xj| + |yi - yj|` ，其中 `|val|` 表示 `val` 的绝对值。  
请你返回将所有点连接的最小总费用。只有任意两点之间 **有且仅有** 一条简单路径时，才认为所有点都已连接。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2020/08/26/d.png)  
Input:  
```
points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
```
Output:
```
20
```
![](https://assets.leetcode.com/uploads/2020/08/26/c.png)  
我们可以按照上图所示连接所有点得到最小总费用，总费用为 20 。  
注意到任意两个点之间只有唯一一条路径互相到达。
### Example 2
Input:  
```
points = [[3,12],[-2,5],[-4,1]]
```
Output:
```
18
```
### Example 3
Input:  
```
points = [[0,0],[1,1],[1,0],[-1,1]]
```
Output:
```
4
```
### Example 4
Input:  
```
points = [[-1000000,-1000000],[1000000,1000000]]
```
Output:
```
4000000
```
## Hint
- 1 <= points.length <= 1000
- -10^6 <= xi, yi <= 10^6
- 所有点 (xi, yi) 两两不同。

