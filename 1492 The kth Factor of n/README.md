# [1492. n 的第 k 个因子](https://leetcode-cn.com/problems/the-kth-factor-of-n/)
## Description
给你两个正整数 `n` 和 `k` 。  
如果正整数 `i` 满足 `n % i == 0` ，那么我们就说正整数 `i` 是整数 `n` 的因子。  
考虑整数 `n` 的所有因子，将它们 **升序排列** 。请你返回第 `k` 个因子。如果 `n` 的因子数少于 `k` ，请你返回 `-1` 。  
## Example
### Example 1
Input:  
```
n = 12, k = 3
```
Output:
```
3
```
因子列表包括 [1, 2, 3, 4, 6, 12]，第 3 个因子是 3 。
### Example 2
Input:  
```
n = 7, k = 2
```
Output:
```
7
```
因子列表包括 [1, 7] ，第 2 个因子是 7 。
### Example 3
Input:
```
n = 4, k = 4
```
Output:
```
-1
```
因子列表包括 [1, 2, 4] ，只有 3 个因子，所以我们应该返回 -1 。
### Example 4
Input:
```
n = 1, k = 1
```
Output:
```
1
```
因子列表包括 [1] ，第 1 个因子为 1 。
### Example 5
Input:
```
n = 1000, k = 3
```
Output:
```
4
```
因子列表包括 [1, 2, 4, 5, 8, 10, 20, 25, 40, 50, 100, 125, 200, 250, 500, 1000] 。
## Hint
- 1 <= k <= n <= 1000
