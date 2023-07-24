# [50. Pow(x, n)](https://leetcode.cn/problems/powx-n/)
## Description
实现 `pow(x, n)` ，即计算 `x` 的整数 `n` 次幂函数（即，`x^n` ）。
## Example
### Example 1
Input:  
```
x = 2.00000, n = 10
```
Output:
```
1024.00000
```
### Example 2
Input:  
```
x = 2.10000, n = 3
```
Output:
```
9.26100
```
### Example 3
Input:  
```
x = 2.00000, n = -2
```
Output:
```
0.25000
```
2^-2 = 1/2^2 = 1/4 = 0.25
## Hint
- -100.0 < x < 100.0
- -2^31 <= n <= 2^31-1
- n 是一个整数
- 要么 x 不为零，要么 n > 0 。
- -10^4 <= xn <= 10^4
