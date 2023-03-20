# [116. 填充每个节点的下一个右侧节点指针](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/)
## Description
给定一个 **完美二叉树** ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：  
```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```
填充它的每个 `next` 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 `next` 指针设置为 `NULL`。  
初始状态下，所有 `next` 指针都被设置为 `NULL`。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2019/02/14/116_sample.png)  
Input:  
```
root = [1,2,3,4,5,6,7]
```
Output:
```
[1,#,2,3,#,4,5,6,7,#]
```
给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。  
### Example 2
Input:  
```
root = []
```
Output:
```
[]
```
## Hint
- 树中节点的数量在 [0, 2^12 - 1] 范围内
- -1000 <= node.val <= 1000

