# [565. 数组嵌套](https://leetcode-cn.com/problems/array-nesting/)
## Description
索引从`0`开始长度为`N`的数组`A`，包含`0`到`N - 1`的所有整数。找到最大的集合`S`并返回其大小，其中 `S[i] = {A[i], A[A[i]], A[A[A[i]]], ... }` 且遵守以下的规则。  
假设选择索引为`i`的元素`A[i]`为`S`的第一个元素，`S`的下一个元素应该是`A[A[i]]`，之后是`A[A[A[i]]]...` 以此类推，不断添加直到`S`出现重复的元素。  
## Example
### Example 1
Input:  
```
A = [5,4,0,3,1,6,2]
```
Output:
```
4
```
A[0] = 5, A[1] = 4, A[2] = 0, A[3] = 3, A[4] = 1, A[5] = 6, A[6] = 2.  
其中一种最长的 S[K]:  
S[0] = {A[0], A[5], A[6], A[2]} = {5, 6, 2, 0}  
## Hint
- N是[1, 20,000]之间的整数。
- A中不含有重复的元素。
- A中的元素大小在[0, N-1]之间。