# [958. 二叉树的完全性检验](https://leetcode.cn/problems/check-completeness-of-a-binary-tree/)
## Description
给定一个二叉树的 `root` ，确定它是否是一个 **完全二叉树** 。  
在一个 `完全二叉树` 中，除了最后一个关卡外，所有关卡都是完全被填满的，并且最后一个关卡中的所有节点都是尽可能靠左的。它可以包含 `1` 到 `2^h` 节点之间的最后一级 `h` 。   
## Example
### Example 1
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/complete-binary-tree-1.png)  
Input:  
```
root = [1,2,3,4,5,6]
```
Output:
```
true
```
最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。
### Example 2
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/complete-binary-tree-2.png)  
Input:  
```
root = [1,2,3,4,5,null,7]
```
Output:
```
false
```
值为 7 的结点没有尽可能靠向左侧。
## Hint
- 树的结点数在范围  [1, 100] 内。
- 1 <= Node.val <= 1000

