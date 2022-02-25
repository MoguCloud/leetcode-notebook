# [1871. 跳跃游戏 VII](https://leetcode-cn.com/problems/jump-game-vii/)
## Description
给你一个下标从 `0` 开始的二进制字符串 `s` 和两个整数 `minJump` 和 `maxJump` 。一开始，你在下标 `0` 处，且该位置的值一定为 `'0'` 。当同时满足如下条件时，你可以从下标 `i` 移动到下标 `j` 处：  
-    i + minJump <= j <= min(i + maxJump, s.length - 1) 且
-    s[j] == '0'.  
如果你可以到达 `s` 的下标 `s.length - 1` 处，请你返回 `true` ，否则返回 `false` 。
## Example
### Example 1
Input:  
```
s = "011010", minJump = 2, maxJump = 3
```
Output:
```
true
```
第一步，从下标 0 移动到下标 3 。  
第二步，从下标 3 移动到下标 5 。  
### Example 2
Input:  
```
s = "01101110", minJump = 2, maxJump = 3
```
Output:
```
false
```
## Hint
- 2 <= s.length <= 105
- s[i] 要么是 '0' ，要么是 '1'
- s[0] == '0'
- 1 <= minJump <= maxJump < s.length
