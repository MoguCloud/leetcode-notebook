# [910. 最小差值 II](https://leetcode-cn.com/problems/smallest-range-ii/)
## Description
给你一个整数数组 `A`，对于每个整数 `A[i]`，可以选择 `x = -K` 或是 `x = K` （K 总是非负整数），并将 `x` 加到 `A[i]` 中。  
在此过程之后，得到数组 `B`。  
返回 `B` 的最大值和 `B` 的最小值之间可能存在的最小差值。  
## Example
### Example 1
Input:  
```
A = [1], K = 0
```
Output:
```
0
```
B = [1]
### Example 2
Input:  
```
A = [0,10], K = 2
```
Output:
```
6
```
B = [2,8]
### Example 3
Input:  
```
A = [1,3,6], K = 3
```
Output:
```
3
```
B = [4,6,3]
## Hint
- 1 <= A.length <= 10000
- 0 <= A[i] <= 10000
- 0 <= K <= 10000

