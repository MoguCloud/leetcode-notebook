# [1721. 交换链表中的节点](https://leetcode.cn/problems/swapping-nodes-in-a-linked-list/)
## Description
给你链表的头节点 `head` 和一个整数 `k` 。  
**交换** 链表正数第 `k` 个节点和倒数第 `k` 个节点的值后，返回链表的头节点（链表 **从 1 开始索引**）。
## Example
### Example 1
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/01/10/linked1.jpg)  
Input:  
```
head = [1,2,3,4,5], k = 2
```
Output:
```
[1,4,3,2,5]
```
### Example 2
Input:  
```
head = [7,9,6,6,7,8,3,0,9,5], k = 5
```
Output:
```
[7,9,6,6,8,7,3,0,9,5]
```
### Example 3
Input:  
```
head = [1], k = 1
```
Output:
```
[1]
```
### Example 4
Input:  
```
head = [1,2], k = 1
```
Output:
```
[2,1]
```
### Example 5
Input:  
```
head = [1,2,3], k = 2
```
Output:
```
[1,2,3]
```
## Hint
- 链表中节点的数目是 n
- 1 <= k <= n <= 10^5
- 0 <= Node.val <= 100

