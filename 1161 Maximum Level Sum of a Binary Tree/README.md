# [1161. 最大层内元素和](https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree/)
## Description
给你一个二叉树的根节点 `root`。设根节点位于二叉树的第 `1` 层，而根节点的子节点位于第 `2` 层，依此类推。  
请你找出层内元素之和 **最大** 的那几层（可能只有一层）的层号，并返回其中 **最小** 的那个。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2019/05/03/capture.JPG)  
Input:  
```
root = [1,7,0,7,-8,null,null]
```
Output:
```
2
```
第 1 层各元素之和为 1，  
第 2 层各元素之和为 7 + 0 = 7，  
第 3 层各元素之和为 7 + -8 = -1，  
所以我们返回第 2 层的层号，它的层内元素之和最大。  
### Example 2
Input:  
```
root = [989,null,10250,98693,-89388,null,null,null,-32127]
```
Output:
```
2
```
## Hint
- 树中的节点数介于 1 和 10^4 之间
- -10^5 <= node.val <= 10^5

