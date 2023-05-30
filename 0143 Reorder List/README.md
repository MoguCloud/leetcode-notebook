# [143. 重排链表](https://leetcode.cn/problems/reorder-list/)  
## Description
给定一个单链表 `L` 的头节点 `head` ，单链表 `L` 表示为：  
```
L0 → L1 → … → Ln - 1 → Ln
```
请将其重新排列后变为：  
```
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
```
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。  
## Example
### Example 1
![](https://pic.leetcode-cn.com/1626420311-PkUiGI-image.png)   
Input:    
```
head = [1,2,3,4]
```
Output:
```
[1,4,2,3]
```
### Example 2
![](https://pic.leetcode-cn.com/1626420320-YUiulT-image.png)   
Input:  
```
head = [1,2,3,4,5]
```
Output:
```
[1,5,2,4,3]
```
## Hint
- 链表的长度范围为 [1, 5 * 10^4]
- 1 <= node.val <= 1000

