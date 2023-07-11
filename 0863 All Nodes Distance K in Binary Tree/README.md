# [863. 二叉树中所有距离为 K 的结点](https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/)
## Description
给定一个二叉树（具有根结点 `root`）， 一个目标结点 `target` ，和一个整数值 `k` 。  
返回到目标结点 `target` 距离为 `k` 的所有结点的值的列表。 答案可以以 **任何顺序** 返回。
## Example
### Example 1
![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png)  
Input:  
```
root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
```
Output:
```
[7,4,1]
```
所求结点为与目标结点（值为 5）距离为 2 的结点，值分别为 7，4，以及 1
### Example 2
Input:  
```
root = [1], target = 1, k = 3
```
Output:
```
[]
```
## Hint
- 节点数在 [1, 500] 范围内
- 0 <= Node.val <= 500
- Node.val 中所有值 不同
- 目标结点 target 是树上的结点。
- 0 <= k <= 1000
