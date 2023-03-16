# [106. 从中序与后序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
## Description
给定两个整数数组 `inorder` 和 `postorder` ，其中 `inorder` 是二叉树的中序遍历， `postorder` 是同一棵树的后序遍历，请你构造并返回这颗 **二叉树** 。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)  
Input:  
```
inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
```
Output:
```
[3,9,20,null,null,15,7]
```
### Example 2
Input:  
```
inorder = [-1], postorder = [-1]
```
Output:
```
[-1]
```
## Hint
- 1 <= inorder.length <= 3000
- postorder.length == inorder.length
- -3000 <= inorder[i], postorder[i] <= 3000
- inorder 和 postorder 都由 不同 的值组成
- postorder 中每一个值都在 inorder 中
- inorder 保证是树的中序遍历
- postorder 保证是树的后序遍历