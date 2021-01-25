# [117. 填充每个节点的下一个右侧节点指针 II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)
## Description
给定一个二叉树  
```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```
填充它的每个 `next` 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 `next` 指针设置为 `NULL`。  
初始状态下，所有 `next` 指针都被设置为 `NULL`。  
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2019/02/15/117_sample.png)
Input:  
```
root = [1,2,3,4,5,null,7]
```
Output:
```
[1,#,2,3,#,4,5,7,#]
```
给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
## Hint
- 1 <= n <= 231 - 1
