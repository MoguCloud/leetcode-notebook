# [19. 删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)
## Description
给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。  
**进阶**：你能尝试使用一趟扫描实现吗？
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)  
Input:  
```
head = [1,2,3,4,5], n = 2
```
Output:
```
[1,2,3,5]
```
### Example 2
Input:  
```
head = [1], n = 1
```
Output:
```
[]
```
### Example 3
Input:  
```
head = [1,2], n = 1
```
Output:
```
1
```
## Hint
- 链表中结点的数目为 sz
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz
