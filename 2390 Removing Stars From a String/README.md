# [2390. 从字符串中移除星号](https://leetcode.cn/problems/removing-stars-from-a-string/)
## Description
给你一个包含若干星号 `*` 的字符串 `s` 。  
在一步操作中，你可以：  
- 选中 `s` 中的一个星号。
- 移除星号 **左侧** 最近的那个 **非星号** 字符，并移除该星号自身。  
返回移除 **所有** 星号之后的字符串。  
注意：  
- 生成的输入保证总是可以执行题面中描述的操作。
- 可以证明结果字符串是唯一的。
## Example
### Example 1
Input:  
```
s = "leet**cod*e"
```
Output:
```
"lecoe"
```
### Example 2
Input:  
```
s = "erase*****"
```
Output:
```
""
```
整个字符串都会被移除，所以返回空字符串。
## Hint
- 1 <= s.length <= 10^5
- s 由小写英文字母和星号 * 组成
- s 可以执行上述操作

