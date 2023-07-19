# [435. 无重叠区间](https://leetcode.cn/problems/non-overlapping-intervals/)
## Description
给定一个区间的集合 `intervals` ，其中 `intervals[i] = [starti, endi]` 。返回 *需要移除区间的最小数量，使剩余区间互不重叠* 。
## Example
### Example 1
Input:  
```
intervals = [[1,2],[2,3],[3,4],[1,3]]
```
Output:
```
1
```
移除 [1,3] 后，剩下的区间没有重叠。
### Example 2
Input:  
```
intervals = [ [1,2], [1,2], [1,2] ]
```
Output:
```
2
```
你需要移除两个 [1,2] 来使剩下的区间没有重叠。
### Example 3
Input:  
```
intervals = [ [1,2], [2,3] ]
```
Output:
```
0
```
你不需要移除任何区间，因为它们已经是无重叠的了。
## Hint
- 1 <= intervals.length <= 10^5
- intervals[i].length == 2
- -5 * 10^4 <= starti < endi <= 5 * 10^4
