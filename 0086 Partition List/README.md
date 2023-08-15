# [86. 分隔链表](https://leetcode.cn/problems/partition-list/description/)
## Description
给你一个链表的头节点 `head` 和一个特定值 `x` ，请你对链表进行分隔，使得所有 **小于** `x` 的节点都出现在 **大于或等于** `x` 的节点之前。  
你应当 **保留** 两个分区中每个节点的初始相对位置。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2021/01/04/partition.jpg)  
Input:  
```
head = [1,4,3,2,5,2], x = 3
```
Output:
```
[1,2,2,4,3,5]
```
### Example 2
Input:  
```
head = [2,1], x = 2
```
Output:
```
[1,2]
```
## Hint
- 链表中节点的数目在范围 [0, 200] 内
- -100 <= Node.val <= 100
- -200 <= x <= 200

