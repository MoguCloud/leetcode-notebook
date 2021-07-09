# [1638. 统计只差一个字符的子串数目](https://leetcode-cn.com/problems/count-substrings-that-differ-by-one-character/)
## Description
给你两个字符串 `s` 和 `t` ，请你找出 `s` 中的非空子串的数目，这些子串满足替换 **一个不同字符** 以后，是 `t` 串的子串。换言之，请你找到 `s` 和 `t` 串中 **恰好** 只有一个字符不同的子字符串对的数目。  
比方说， "**compute**r" 和 "**computa**tion" 加粗部分只有一个字符不同： `'e'/'a'` ，所以这一对子字符串会给答案加 1 。  
请你返回满足上述条件的不同子字符串对数目。  
一个 **子字符串** 是一个字符串中连续的字符。
## Example
### Example 1
Input:  
```
s = "aba", t = "baba"
```
Output:
```
6
```
以下为只相差 1 个字符的 s 和 t 串的子字符串对：  
("**a**ba", "**b**aba")  
("**a**ba", "ba**b**a")  
("ab**a**", "**b**aba")  
("ab**a**", "ba**b**a")  
("a**b**a", "b**a**ba")  
("a**b**a", "bab**a**")  
加粗部分分别表示 s 和 t 串选出来的子字符串。
### Example 2
Input:  
```
s = "ab", t = "bb"
```
Output:
```
3
```
以下为只相差 1 个字符的 s 和 t 串的子字符串对：  
("**a**b", "**b**b")  
("**a**b", "b**b**")  
("**ab**", "**bb**")  
加粗部分分别表示 s 和 t 串选出来的子字符串。
### Example 3
Input:  
```
s = "a", t = "a"
```
Output:
```
0
```
### Example 4
Input:  
```
s = "abe", t = "bbc"
```
Output:
```
10
```
## Hint
- 1 <= s.length, t.length <= 100
- s 和 t 都只包含小写英文字母。

