# [1209. 删除字符串中的所有相邻重复项 II](https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii/)
## Description
给你一个字符串 `s`，「`k` 倍重复项删除操作」将会从 `s` 中选择 `k` 个相邻且相等的字母，并删除它们，使被删去的字符串的左侧和右侧连在一起。  
你需要对 `s` 重复进行无限次这样的删除操作，直到无法继续为止。  
在执行完所有删除操作后，返回最终得到的字符串。  
本题答案保证唯一。  
## Example
### Example 1
Input:  
```
s = "abcd", k = 2
```
Output:
```
"abcd"
```
没有要删除的内容。
### Example 2
Input:  
```
s = "deeedbbcccbdaa", k = 3
```
Output:
```
"aa"
```
先删除 "eee" 和 "ccc"，得到 "ddbbbdaa"  
再删除 "bbb"，得到 "dddaa"  
最后删除 "ddd"，得到 "aa"  
### Example 3
Input:  
```
s = "pbbcggttciiippooaais", k = 2
```
Output:
```
"ps"
```
## Hint
- 1 <= s.length <= 10^5
- 2 <= k <= 10^4
- s 中只含有小写英文字母。
