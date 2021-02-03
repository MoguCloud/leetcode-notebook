# [1410. HTML 实体解析器](https://leetcode-cn.com/problems/html-entity-parser/)
## Description
「HTML 实体解析器」 是一种特殊的解析器，它将 HTML 代码作为输入，并用字符本身替换掉所有这些特殊的字符实体。  
HTML 里这些特殊字符和它们对应的字符实体包括：  
- 双引号：字符实体为 `&quot;` ，对应的字符是 `"` 。
- 单引号：字符实体为 `&apos;` ，对应的字符是 `'` 。
- 与符号：字符实体为 `&amp;` ，对应对的字符是 `&` 。
- 大于号：字符实体为 `&gt;` ，对应的字符是 `>` 。
- 小于号：字符实体为 `&lt;` ，对应的字符是 `<` 。
- 斜线号：字符实体为 `&frasl;` ，对应的字符是 `/` 。
给你输入字符串 `text` ，请你实现一个 HTML 实体解析器，返回解析器解析后的结果。
## Example
### Example 1
Input:  
```
text = "&amp; is an HTML entity but &ambassador; is not."
```
Output:
```
"& is an HTML entity but &ambassador; is not."
```
解析器把字符实体 &amp; 用 & 替换
### Example 2
Input:  
```
text = "and I quote: &quot;...&quot;"
```
Output:
```
"and I quote: \"...\""
```
### Example 3
Input:  
```
text = "Stay home! Practice on Leetcode :)"
```
Output:
```
"Stay home! Practice on Leetcode :)"
```
### Example 4
Input:  
```
text = "x &gt; y &amp;&amp; x &lt; y is always false"
```
Output:
```
"x > y && x < y is always false"
```
### Example 5
Input:  
```
text = "leetcode.com&frasl;problemset&frasl;all"
```
Output:
```
"leetcode.com/problemset/all"
```
## Hint
- 1 <= text.length <= 10^5
- 字符串可能包含 256 个ASCII 字符中的任意字符。
