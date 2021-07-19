# [979. 在二叉树中分配硬币](https://leetcode-cn.com/problems/distribute-coins-in-binary-tree/)
## Description
给定一个有 `N` 个结点的二叉树的根结点 `root`，树中的每个结点上都对应有 `node.val` 枚硬币，并且总共有 `N` 枚硬币。  
在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点。(移动可以是从父结点到子结点，或者从子结点移动到父结点。)。  
返回使每个结点上只有一枚硬币所需的移动次数。  
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2019/01/18/tree1.png)  
Input:  
```
[3,0,0]
```
Output:
```
2
```
从树的根结点开始，我们将一枚硬币移到它的左子结点上，一枚硬币移到它的右子结点上。
### Example 2
![](https://assets.leetcode.com/uploads/2019/01/18/tree2.png)  
Input:  
```
[0,3,0]
```
Output:
```
3
```
从根结点的左子结点开始，我们将两枚硬币移到根结点上 [移动两次]。然后，我们把一枚硬币从根结点移到右子结点上。
### Example 3
![](https://assets.leetcode.com/uploads/2019/01/18/tree3.png)  
Input:  
```
[1,0,2]
```
Output:
```
2
```
### Example 4
![](https://assets.leetcode.com/uploads/2019/01/18/tree4.png)  
Input:  
```
[1,0,0,null,3]
```
Output:
```
4
```
## Hint
- 1<= N <= 100
- 0 <= node.val <= N
