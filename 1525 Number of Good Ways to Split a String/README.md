# [1525. 字符串的好分割数目](https://leetcode-cn.com/problems/number-of-good-ways-to-split-a-string/)
## Description
给你一个字符串 `s` ，一个分割被称为 **「好分割」** 当它满足：将 `s` 分割成 `2` 个字符串 `p` 和 `q` ，它们连接起来等于 `s` 且 `p` 和 `q` 中不同字符的数目相同。  
请你返回 `s` 中好分割的数目。
## Example
### Example 1
Input:  
```
s = "aacaba"
```
Output:
```
2
```
总共有 5 种分割字符串 "aacaba" 的方法，其中 2 种是好分割。  
("a", "acaba") 左边字符串和右边字符串分别包含 1 个和 3 个不同的字符。  
("aa", "caba") 左边字符串和右边字符串分别包含 1 个和 3 个不同的字符。  
("aac", "aba") 左边字符串和右边字符串分别包含 2 个和 2 个不同的字符。这是一个好分割。  
("aaca", "ba") 左边字符串和右边字符串分别包含 2 个和 2 个不同的字符。这是一个好分割。  
("aacab", "a") 左边字符串和右边字符串分别包含 3 个和 1 个不同的字符。  
### Example 2
Input:  
```
s = "abcd"
```
Output:
```
1
```
好分割为将字符串分割成 ("ab", "cd") 。
### Example 3
Input:  
```
s = "aaaaa"
```
Output:
```
4
```
所有分割都是好分割。
### Example 4
Input:  
```
s = "acbadbaada"
```
Output:
```
2
```
## Hint
- s 只包含小写英文字母。
- 1 <= s.length <= 10^5

