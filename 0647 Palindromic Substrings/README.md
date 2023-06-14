# [647. 回文子串](https://leetcode.cn/problems/palindromic-substrings/)  
## Description
给你一个字符串 `s` ，请你统计并返回这个字符串中 **回文子串** 的数目。   
**回文字符串** 是正着读和倒过来读一样的字符串。    
**子字符串** 是字符串中的由连续字符组成的一个序列。  
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。  
## Example
### Example 1
Input:  
```
s = "abc"
```
Output:
```
3
```
三个回文子串: "a", "b", "c"
### Example 2
Input:  
```
s = "aaa"
```
Output:
```
6
```
6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
## Hint
- 1 <= s.length <= 1000
- s 由小写英文字母组成
