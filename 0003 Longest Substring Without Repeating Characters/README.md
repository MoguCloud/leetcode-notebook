# [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)  
## Description
给定一个字符串 `s` ，请你找出其中不含有重复字符的 **最长子串** 的长度。   
## Example
### Example 1
Input:  
```
s = "abcabcbb"
```
Output:
```
3
```
因为无重复字符的最长子串是 "abc"，所以其长度为 3。

### Example 2
Input:  
```
s = "bbbbb"
```
Output:
```
1
```
因为无重复字符的最长子串是 "b"，所以其长度为 1。
### Example 3
Input:  
```
s = "pwwkew"
```
Output:
```
3
```
因为无重复字符的最长子串是 "wke"，所以其长度为 3。  
请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
## Hint
- 0 <= s.length <= 5 * 10^4
- s 由英文字母、数字、符号和空格组成

