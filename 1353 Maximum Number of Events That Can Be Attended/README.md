# [1353. 最多可以参加的会议数目](https://leetcode-cn.com/problems/maximum-number-of-events-that-can-be-attended/)
## Description
给你一个数组 `events`，其中 `events[i] = [startDayi, endDayi]` ，表示会议 `i` 开始于 `startDayi` ，结束于 `endDayi` 。  
你可以在满足 `startDayi <= d <= endDayi` 中的任意一天 `d` 参加会议 `i` 。注意，一天只能参加一个会议。  
请你返回你可以参加的 **最大** 会议数目。  
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2020/02/05/e1.png)
Input:  
```
events = [[1,2],[2,3],[3,4]]
```
Output:
```
3
```
你可以参加所有的三个会议。  
安排会议的一种方案如上图。  
第 1 天参加第一个会议。  
第 2 天参加第二个会议。  
第 3 天参加第三个会议。  
### Example 2
Input:  
```
events = [[1,2],[2,3],[3,4],[1,2]]
```
Output:
```
4
```
### Example 3
Input:  
```
events = [[1,4],[4,4],[2,2],[3,4],[1,1]]
```
Output:
```
4
```
### Example 4
Input:  
```
events = [[1,100000]]
```
Output:
```
1
```
### Example 5
Input:  
```
events = [[1,1],[1,2],[1,3],[1,4],[1,5],[1,6],[1,7]]
```
Output:
```
7
```
## Hint
- 1 <= events.length <= 10^5
- events[i].length == 2
- 1 <= events[i][0] <= events[i][1] <= 10^5
