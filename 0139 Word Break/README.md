# [139. 单词拆分](https://leetcode.cn/problems/word-break/)
## Description
给你一个字符串 `s` 和一个字符串列表 `wordDict` 作为字典。请你判断是否可以利用字典中出现的单词拼接出 `s` 。  
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。  
## Example
### Example 1
Input:  
```
s = "leetcode", wordDict = ["leet", "code"]
```
Output:
```
true
```
返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
### Example 2
Input:  
```
s = "applepenapple", wordDict = ["apple", "pen"]
```
Output:
```
true
```
返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。  
注意，你可以重复使用字典中的单词。
### Example 3
Input:  
```
s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
```
Output:
```
false
```
## Hint
- 1 <= s.length <= 300
- 1 <= wordDict.length <= 1000
- 1 <= wordDict[i].length <= 20
- s 和 wordDict[i] 仅有小写英文字母组成
- wordDict 中的所有字符串 互不相同
