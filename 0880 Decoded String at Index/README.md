# [880. 索引处的解码字符串](https://leetcode.cn/problems/decoded-string-at-index/)  
## Description
给定一个编码字符串 `S`。请你找出 **解码字符串** 并将其写入磁带。解码时，从编码字符串中 **每次读取一个字符** ，并采取以下步骤：  
- 如果所读的字符是字母，则将该字母写在磁带上。  
- 如果所读的字符是数字（例如 `d`），则整个当前磁带总共会被重复写 `d-1` 次。  



现在，对于给定的编码字符串 `S` 和索引 `K`，查找并返回解码字符串中的第 `K` 个字母。
## Example
### Example 1
Input:  
```
S = "leet2code3", K = 10
```
Output:
```
"o"
```
解码后的字符串为 "leetleetcodeleetleetcodeleetleetcode"。  
字符串中的第 10 个字母是 "o"。
### Example 2
Input:  
```
S = "ha22", K = 5
```
Output:
```
"h"
```
解码后的字符串为 "hahahaha"。第 5 个字母是 "h"。
### Example 3
Input:  
```
S = "a2345678999999999999999", K = 1
```
Output:
```
"a"
```
解码后的字符串为 "a" 重复 8301530446056247680 次。第 1 个字母是 "a"。
## Hint
- 2 <= S.length <= 100
- S 只包含小写字母与数字 2 到 9 。
- S 以字母开头。
- 1 <= K <= 10^9
- 题目保证 K 小于或等于解码字符串的长度。
- 解码后的字符串保证少于 2^63 个字母。
