# [1191. K 次串联后最大子数组之和](https://leetcode-cn.com/problems/k-concatenation-maximum-sum/)
## Description
给你一个整数数组 `arr` 和一个整数 `k`。  
首先，我们要对该数组进行修改，即把原数组 `arr` 重复 `k` 次。  

> 举个例子，如果 arr = [1, 2] 且 k = 3，那么修改后的数组就是 [1, 2, 1, 2, 1, 2]。  

然后，请你返回修改后的数组中的最大的子数组之和。  
注意，子数组长度可以是 `0`，在这种情况下它的总和也是 `0`。  
由于 **结果可能会很大**，所以需要 **模（mod）** `10^9 + 7` 后再返回。  
## Example
### Example 1
Input:  
```
arr = [1,2], k = 3
```
Output:
```
9
```
### Example 2
Input:  
```
arr = [1,-2,1], k = 5
```
Output:
```
2
```
### Example 3
Input:  
```
arr = [-1,-2], k = 7
```
Output:
```
0
```
## Hint
- 1 <= arr.length <= 10^5
- 1 <= k <= 10^5
- -10^4 <= arr[i] <= 10^4

