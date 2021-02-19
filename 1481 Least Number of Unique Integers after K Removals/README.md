# [1481. 不同整数的最少数目](https://leetcode-cn.com/problems/least-number-of-unique-integers-after-k-removals/)
## Description
给你一个整数数组 `arr` 和一个整数 `k` 。现需要从数组中恰好移除 `k` 个元素，请找出移除后数组中不同整数的最少数目。
## Example
### Example 1
Input:  
```
arr = [5,5,4], k = 1
```
Output:
```
1
```
移除 1 个 4 ，数组中只剩下 5 一种整数。
### Example 2
Input:  
```
arr = [4,3,1,1,3,3,2], k = 3
```
Output:
```
2
```
先移除 4、2 ，然后再移除两个 1 中的任意 1 个或者三个 3 中的任意 1 个，最后剩下 1 和 3 两种整数。
## Hint
- 1 <= arr.length <= 10^5
- 1 <= arr[i] <= 10^9
- 0 <= k <= arr.length
