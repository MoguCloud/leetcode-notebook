# [1461. 检查一个字符串是否包含所有长度为 K 的二进制子串](https://leetcode-cn.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/)
## Description
给你一个二进制字符串 `s` 和一个整数 `k` 。  
如果所有长度为 `k` 的二进制字符串都是 `s` 的子串，请返回 True ，否则请返回 False 。  
## Example
### Example 1
Input:  
```
s = "00110110", k = 2
```
Output:  
```
true
```
长度为 2 的二进制串包括 "00"，"01"，"10" 和 "11"。它们分别是 s 中下标为 0，1，3，2 开始的长度为 2 的子串。  
### Example 2
Input:  
```
s = "00110", k = 2
```
Output:
```
true
```
### Example 3
Input:  
```
s = "0110", k = 1
```
Output:
```
true
```
长度为 1 的二进制串包括 "0" 和 "1"，显然它们都是 s 的子串。
### Example 4
Input:  
```
s = "0110", k = 2
```
Output:
```
false
```
长度为 2 的二进制串 "00" 没有出现在 s 中。  
### Example 5
Input:  
```
s = "0000000001011100", k = 4
```
Output:
```
false
```
## Hint
- 1 <= s.length <= 5 * 10^5
- s 中只含 0 和 1 。
- 1 <= k <= 20
