# [397. 整数替换](https://leetcode-cn.com/problems/integer-replacement/)
## Description
给定一个正整数 `n` ，你可以做如下操作：
1. 如果 `n` 是偶数，则用 `n / 2`替换 `n` 。
2. 如果 `n` 是奇数，则可以用 `n + 1`或`n - 1`替换 `n` 。
`n` 变为 `1` 所需的最小替换次数是多少？
## Example
### Example 1
Input:  
```
n = 8
```
Output:
```
3
```
8 -> 4 -> 2 -> 1
### Example 2
Input:  
```
n = 7
```
Output:
```
4
```
7 -> 8 -> 4 -> 2 -> 1
或 7 -> 6 -> 3 -> 2 -> 1
### Example 3
Input:  
```
n = 4
```
Output:
```
2
```
## Hint
- 1 <= n <= 231 - 1
