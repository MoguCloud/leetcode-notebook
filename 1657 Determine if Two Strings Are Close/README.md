# [1657. 确定两个字符串是否接近](https://leetcode-cn.com/problems/determine-if-two-strings-are-close/)
## Description
如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 **接近** ：  
- 操作 1：交换任意两个 **现有** 字符。  
 - 例如，a**b**cde -> aecd**b**  
- 操作 2：将一个 **现有** 字符的每次出现转换为另一个 **现有** 字符，并对另一个字符执行相同的操作。  
 - 例如，**aa**c**abb** -> **bb**c**baa**（所有 `a` 转化为 `b` ，而所有 `b` 转换为 `a` ）  
你可以根据需要对任意一个字符串多次使用这两种操作。  
给你两个字符串，`word1` 和 `word2` 。如果 `word1` 和 `word2` **接近** ，就返回 `true` ；否则，返回 `false` 。  

## Example
### Example 1
Input:  
```
word1 = "abc", word2 = "bca"
```
Output:
```
true
```
2 次操作从 word1 获得 word2 。  
执行操作 1："a**b**c" -> "ac**b**"  
执行操作 1："**a**c**b**" -> "**b**c**a**"  
### Example 2
Input:  
```
word1 = "a", word2 = "aa"
```
Output:
```
false
```
不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
### Example 3
Input:
```
word1 = "cabbba", word2 = "abbccc"
```
Output:
```
true
```
3 次操作从 word1 获得 word2 。  
执行操作 1："ca**b**bb**a**" -> "ca**a**bb**b**"  
执行操作 2："**c**aa**bbb**" -> "**b**aa**ccc**"  
执行操作 2："**baa**ccc" -> "**abb**ccc"  
### Example 4
Input:
```
word1 = "cabbba", word2 = "aabbss"
```
Output:
```
false
```
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
## Hint
- 1 <= word1.length, word2.length <= 105
- word1 和 word2 仅包含小写英文字母

