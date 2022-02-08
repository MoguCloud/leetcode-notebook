# [916. 单词子集](https://leetcode-cn.com/problems/word-subsets/)
## Description
给你两个字符串数组 `words1` 和 `words2`。  
现在，如果 `b` 中的每个字母都出现在 `a` 中，**包括重复出现的字母**，那么称字符串 `b` 是字符串 `a` 的 **子集** 。  
- 例如，"wrr" 是 "warrior" 的子集，但不是 "world" 的子集。  
如果对 `words2` 中的每一个单词 `b`，`b` 都是 `a` 的子集，那么我们称 `words1` 中的单词 `a` 是 **通用单词** 。  
以数组形式返回 `words1` 中所有的通用单词。你可以按 **任意顺序** 返回答案。  
## Example
### Example 1
Input:  
```
words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
```
Output:
```
["facebook","google","leetcode"]
```
### Example 2
Input:  
```
words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
```
Output:
```
["apple","google","leetcode"]
```
### Example 3
Input:  
```
words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","oo"]
```
Output:
```
["facebook","google"]
```
### Example 4
Input:  
```
words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["lo","eo"]
```
Output:
```
["google","leetcode"]
```
### Example 5
Input:  
```
words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["ec","oc","ceo"]
```
Output:
```
["facebook","leetcode"]
```
## Hint
- 1 <= words1.length, words2.length <= 10^4
- 1 <= words1[i].length, words2[i].length <= 10
- words1[i] 和 words2[i] 仅由小写英文字母组成
- words1 中的所有字符串 **互不相同**
