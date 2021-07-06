# [1072. 按列翻转得到最大值等行数](https://leetcode-cn.com/problems/flip-columns-for-maximum-number-of-equal-rows/)
## Description
给定由若干 `0` 和 `1` 组成的矩阵 `matrix`，从中选出任意数量的列并翻转其上的 **每个** 单元格。翻转后，单元格的值从 `0` 变成 `1`，或者从 `1` 变为 `0` 。  
回经过一些翻转后，行与行之间所有值都相等的最大行数。  
## Example
### Example 1
Input:  
```
[[0,1],[1,1]]
```
Output:
```
1
```
不进行翻转，有 1 行所有值都相等。
### Example 2
Input:  
```
[[0,1],[1,0]]
```
Output:
```
2
```
翻转第一列的值之后，这两行都由相等的值组成。
### Example 3
Input:  
```
[[0,0,0],[0,0,1],[1,1,0]]
```
Output:
```
2
```
翻转前两列的值之后，后两行由相等的值组成。
## Hint
- 1 <= matrix.length <= 300
- 1 <= matrix[i].length <= 300
- 所有 matrix[i].length 都相等
- matrix[i][j] 为 0 或 1
