# [17. 电话号码的字母组合](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)
## Description
给定一个仅包含数字 `2-9` 的字符串，返回所有它能表示的字母组合。答案可以按 **任意顺序** 返回。  
给出数字到字母的映射如下（与电话按键相同）。注意 `1` 不对应任何字母。  
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/11/09/200px-telephone-keypad2svg.png)  
## Example
### Example 1
Input:  
```
digits = "23"
```
Output:
```
["ad","ae","af","bd","be","bf","cd","ce","cf"]
```
### Example 2
Input:  
```
digits = ""
```
Output:
```
[]
```
### Example 3
Input:  
```
digits = "2"
```
Output:
```
["a","b","c"]
```
## Hint
- 0 <= digits.length <= 4
- digits[i] 是范围 ['2', '9'] 的一个数字。
