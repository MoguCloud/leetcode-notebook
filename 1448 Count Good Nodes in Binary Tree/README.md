# [1448. 统计二叉树中好节点的数目](https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree/)
## Description
给你一棵根为 `root` 的二叉树，请你返回二叉树中好节点的数目。  
「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png)  
Input:  
```
root = [3,1,4,3,null,1,5]
```
Output:
```
4
```
图中蓝色节点为好节点。  
根节点 (3) 永远是个好节点。  
节点 4 -> (3,4) 是路径中的最大值。  
节点 5 -> (3,4,5) 是路径中的最大值。  
节点 3 -> (3,1,3) 是路径中的最大值。  
### Example 2
![](https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png)  
Input:  
```
root = [3,3,null,4,2]
```
Output:
```
3
```
节点 2 -> (3, 3, 2) 不是好节点，因为 "3" 比它大。
### Example 3
Input:  
```
root = [1]
```
Output:
```
1
```
根节点是好节点。
## Hint
- 二叉树中节点数目范围是 [1, 10^5] 。
- 每个节点权值的范围是 [-10^4, 10^4] 。
