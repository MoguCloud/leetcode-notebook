# [1405. 最长快乐字符串](https://leetcode-cn.com/problems/longest-happy-string/)
## Description
如果字符串中不含有任何 `'aaa'`，`'bbb'` 或 `'ccc'` 这样的字符串作为子串，那么该字符串就是一个「快乐字符串」。  
给你三个整数 `a`，`b` ，`c`，请你返回 **任意一个** 满足下列全部条件的字符串 `s`：  
- `s` 是一个尽可能长的快乐字符串。
- `s` 中 **最多** 有 `a` 个字母 `'a'`、`b` 个字母 `'b'`、`c` 个字母 `'c'` 。
- `s` 中只含有 `'a'`、`'b'` 、`'c'` 三种字母。
如果不存在这样的字符串 `s` ，请返回一个空字符串 `""`。
## Example
### Example 1
Input:  
```
a = 1, b = 1, c = 7
```
Output:
```
"ccaccbcc"
```
"ccbccacc" 也是一种正确答案。
### Example 2
Input:  
```
a = 2, b = 2, c = 1
```
Output:
```
"aabbc"
```
### Example 3
Input:  
```
a = 7, b = 1, c = 0
```
Output:
```
"aabaa"
```
这是该测试用例的唯一正确答案。
## Hint
- 0 <= a, b, c <= 100
- a + b + c > 0

