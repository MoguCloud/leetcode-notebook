# [1315. 祖父节点值为偶数的节点和](https://leetcode.cn/problems/sum-of-nodes-with-even-valued-grandparent/)
## Description
给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：  
- 该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）  
如果不存在祖父节点值为偶数的节点，那么返回 `0` 。  
## Example
### Example 1
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/01/10/1473_ex1.png)   
Input:  
```
root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
```
Output:
```
18
```
图中红色节点的祖父节点的值为偶数，蓝色节点为这些红色节点的祖父节点。
## Hint
- 树中节点的数目在 1 到 10^4 之间。
- 每个节点的值在 1 到 100 之间。

