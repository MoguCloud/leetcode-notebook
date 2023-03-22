# [567. 字符串的排列](https://leetcode.cn/problems/permutation-in-string/)
## Description
给你两个字符串 `s1` 和 `s2` ，写一个函数来判断 `s2` 是否包含 `s1` 的排列。如果是，返回 `true` ；否则，返回 `false` 。  
换句话说，`s1` 的排列之一是 `s2` 的 **子串** 。
## Example
### Example 1
Input:  
```
s1 = "ab" s2 = "eidbaooo"
```
Output:
```
true
```
s2 包含 s1 的排列之一 ("ba").
### Example 2
Input:  
```
s1= "ab" s2 = "eidboaoo"
```
Output:
```
false
```
## Hint
- 1 <= s1.length, s2.length <= 10^4
- s1 和 s2 仅包含小写字母
