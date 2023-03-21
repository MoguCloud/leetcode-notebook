# [542. 01 矩阵](https://leetcode.cn/problems/01-matrix/)
## Description
给定一个由 `0` 和 `1` 组成的矩阵 `mat` ，请输出一个大小相同的矩阵，其中每一个格子是 `mat` 中对应位置元素到最近的 `0` 的距离。  
两个相邻元素间的距离为 `1` 。  
## Example
### Example 1
![](https://pic.leetcode-cn.com/1626667201-NCWmuP-image.png)  
Input:  
```
mat = [[0,0,0],[0,1,0],[0,0,0]]
```
Output:
```
[[0,0,0],[0,1,0],[0,0,0]]
```
### Example 2
![](https://pic.leetcode-cn.com/1626667205-xFxIeK-image.png)  
Input:  
```
mat = [[0,0,0],[0,1,0],[1,1,1]]
```
Output:
```
[[0,0,0],[0,1,0],[1,2,1]]
```
## Hint
- m == mat.length
- n == mat[i].length
- 1 <= m, n <= 10^4
- 1 <= m * n <= 10^4
- mat[i][j] is either 0 or 1.
- mat 中至少有一个 0 

