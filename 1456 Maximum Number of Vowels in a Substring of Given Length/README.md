# [1456. 定长子串中元音的最大数目](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/)  
## Description
给你字符串 `s` 和整数 `k` 。  
请返回字符串 `s` 中长度为 `k` 的单个子字符串中可能包含的最大元音字母数。  
英文中的 **元音字母** 为（`a`, `e`, `i`, `o`, `u`）。
## Example
### Example 1
Input:  
```
s = "abciiidef", k = 3
```
Output:
```
3
```
子字符串 "iii" 包含 3 个元音字母。
### Example 2
Input:  
```
s = "aeiou", k = 2
```
Output:
```
2
```
任意长度为 2 的子字符串都包含 2 个元音字母。
### Example 3
Input:  
```
s = "leetcode", k = 3
```
Output:
```
2
```
"lee"、"eet" 和 "ode" 都包含 2 个元音字母。
### Example 4
Input:  
```
s = "rhythms", k = 4
```
Output:
```
0
```
字符串 s 中不含任何元音字母。
### Example 5
Input:  
```
s = "tryhard", k = 4
```
Output:
```
1
```
## Hint
- 1 <= s.length <= 10^5
- s 由小写英文字母组成
- 1 <= k <= s.length

