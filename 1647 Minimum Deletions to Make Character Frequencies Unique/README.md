# [1647. 字符频次唯一的最小删除次数](https://leetcode.cn/problems/minimum-deletions-to-make-character-frequencies-unique)
## Description
如果字符串 `s` 中 **不存在** 两个不同字符 **频次** 相同的情况，就称 `s` 是 **优质字符串** 。  
给你一个字符串 `s`，返回使 `s` 成为 **优质字符串** 需要删除的 **最小** 字符数。  
字符串中字符的 **频次** 是该字符在字符串中的出现次数。例如，在字符串 `"aab"` 中，`'a'` 的频次是 `2`，而 `'b'` 的频次是 `1` 。  
## Example
### Example 1
Input:  
```
s = "aab"
```
Output:
```
0
```
s 已经是优质字符串。
### Example 2
Input:  
```
s = "aaabbbcc"
```
Output:
```
2
```
可以删除两个 'b' , 得到优质字符串 "aaabcc" 。  
另一种方式是删除一个 'b' 和一个 'c' ，得到优质字符串 "aaabbc" 。
### Example 3
Input:  
```
s = "ceabaacb"
```
Output:
```
2
```
可以删除两个 'c' 得到优质字符串 "eabaab" 。  
注意，只需要关注结果字符串中仍然存在的字符。（即，频次为 0 的字符会忽略不计。）
## Hint
- 1 <= s.length <= 10^5
- s 仅含小写英文字母

