# [1277. 统计全为 1 的正方形子矩阵](https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/description/)
## Description
给你一个 `m * n` 的矩阵，矩阵中的元素不是 `0 `就是 `1`，请你统计并返回其中完全由 `1` 组成的 **正方形** 子矩阵的个数。
## Example
### Example 1
Input:  
```
matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
```
Output:
```
15
```
边长为 1 的正方形有 10 个。  
边长为 2 的正方形有 4 个。  
边长为 3 的正方形有 1 个。  
正方形的总数 = 10 + 4 + 1 = 15.
### Example 2
Input:  
```
matrix = 
[
  [1,0,1],
  [1,1,0],
  [1,1,0]
]
```
Output:
```
7
```
边长为 1 的正方形有 6 个。   
边长为 2 的正方形有 1 个。  
正方形的总数 = 6 + 1 = 7.  
## Hint
- 1 <= arr.length <= 300
- 1 <= arr[0].length <= 300
- 0 <= arr[i][j] <= 1
