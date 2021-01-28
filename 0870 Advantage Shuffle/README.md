# [870. 优势洗牌](https://leetcode-cn.com/problems/advantage-shuffle/)
## Description
给定两个大小相等的数组 `A` 和 `B`，`A` 相对于 `B` 的**优势**可以用满足 `A[i] > B[i]` 的索引 `i` 的数目来描述。  
返回 `A` 的任意排列，使其相对于 `B` 的**优势**最大化。
## Example
### Example 1
Input:  
```
A = [2,7,11,15], B = [1,10,4,11]
```
Output:
```
[2,11,7,15]
```
### Example 2
Input:  
```
A = [12,24,8,32], B = [13,25,32,11]
```
Output:
```
[24,32,8,12]
```
## Hint
- 1 <= A.length = B.length <= 10000
- 0 <= A[i] <= 10^9
- 0 <= B[i] <= 10^9
