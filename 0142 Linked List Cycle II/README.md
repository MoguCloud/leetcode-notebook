# [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/)
## Description
给定一个链表的头节点  `head` ，返回链表开始入环的第一个节点。 如果链表**无环**，则返回 `null`。  
如果链表中有某个节点，可以通过连续跟踪 `next` 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 `pos` 来表示链表尾连接到链表中的位置**（索引从 0 开始）**。如果 `pos` 是 `-1`，则在该链表中没有环。注意：`pos` **不作为参数进行传递**，仅仅是为了标识链表的实际情况。  
**不允许修改** 链表。  
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)  
Input:  
```
head = [3,2,0,-4], pos = 1
```
Output:
```
返回索引为 1 的链表节点
```
链表中有一个环，其尾部连接到第二个节点。
### Example 2
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png)    
Input:  
```
head = [1,2], pos = 0
```
Output:
```
返回索引为 0 的链表节点
```
链表中有一个环，其尾部连接到第一个节点。
### Example 3
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png)  
Input:  
```
head = [1], pos = -1
```
Output:
```
返回 null
```
链表中没有环。
## Hint
- 链表中节点的数目范围在范围 [0, 10^4] 内
- -10^5 <= Node.val <= 10^5
- pos 的值为 -1 或者链表中的一个有效索引

