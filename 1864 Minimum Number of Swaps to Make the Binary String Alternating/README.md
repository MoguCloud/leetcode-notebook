# [1864. 构成交替字符串需要的最小交换次数](https://leetcode-cn.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/)
## Description
给你一个二进制字符串 `s` ，现需要将其转化为一个 **交替字符串** 。请你计算并返回转化所需的 **最小** 字符交换次数，如果无法完成转化，返回 `-1` 。  
**交替字符串** 是指：相邻字符之间不存在相等情况的字符串。例如，字符串 `"010"` 和 `"1010"` 属于交替字符串，但 `"0100"` 不是。  
任意两个字符都可以进行交换，**不必相邻** 。
## Example
### Example 1
Input:  
```
s = "111000"
```
Output:
```
1
```
交换位置 1 和 4："111000" -> "101010" ，字符串变为交替字符串。
### Example 2
Input:  
```
s = "010"
```
Output:
```
0
```
字符串已经是交替字符串了，不需要交换。
### Example 3
Input:  
```
s = "1110"
```
Output:
```
-1
```
## Hint
- 1 <= s.length <= 1000
- s[i] 的值为 '0' 或 '1'

