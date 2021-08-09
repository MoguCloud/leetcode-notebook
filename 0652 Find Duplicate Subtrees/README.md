# [652. 寻找重复的子树](https://leetcode-cn.com/problems/find-duplicate-subtrees/)
## Description
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意**一棵**的根结点即可。
两棵树重复是指它们具有相同的结构以及相同的结点值。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2020/08/16/e1.jpg)  
Input:  
```
root = [1,2,3,4,null,2,4,null,null,4]
```
Output:
```
[[2,4],[4]]
```
### Example 2
![](https://assets.leetcode.com/uploads/2020/08/16/e2.jpg)  
Input:  
```
root = [2,1,1]
```
Output:
```
[[1]]
```
### Example 3
![](https://assets.leetcode.com/uploads/2020/08/16/e33.jpg)  
Input:  
```
root = [2,2,2,3,null,3,null]
```
Output:
```
[[2,3],[3]]
```
## Hint
- The number of the nodes in the tree will be in the range [1, 10^4]
- -200 <= Node.val <= 200

