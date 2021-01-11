# [1325. 删除给定值的叶子节点](https://leetcode-cn.com/problems/delete-leaves-with-a-given-value/)
## Description
给你一棵以 `root` 为根的二叉树和一个整数 `target` ，请你删除所有值为 `target` 的 叶子节点。  
注意，一旦删除值为 `target` 的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是 `target` ，那么这个节点也应该被删除。  
也就是说，你需要重复此过程直到不能继续删除。   
## Example
### Example 1
![](https://assets.leetcode.com/uploads/2020/01/09/sample_1_1684.png)  
Input:  
```
root = [1,2,3,2,null,2,4], target = 2
```
Output:
```
[1,null,3,null,4]
```
上面左边的图中，绿色节点为叶子节点，且它们的值与 target 相同（同为 2 ），它们会被删除，得到中间的图。  
有一个新的节点变成了叶子节点且它的值与 target 相同，所以将再次进行删除，从而得到最右边的图。  
### Example 2
![](https://assets.leetcode.com/uploads/2020/01/09/sample_2_1684.png)  
Input:  
```
root = [1,3,3,3,2], target = 3
```
Output:
```
[1,3,null,null,2]
```
### Example 3
![](https://assets.leetcode.com/uploads/2020/01/15/sample_3_1684.png)  
Input:  
```
root = [1,2,null,2,null,2], target = 2
```
Output:
```
[1]
```
每一步都删除一个绿色的叶子节点（值为 2）。  
### Example 4
Input:  
```
root = [1,1,1], target = 1
```
Output:
```
[]
```
### Example 5
Input:  
```
root = [1,2,3], target = 1
```
Output:
```
[1,2,3]
```
## Hint
- 1 <= target <= 1000
- 每一棵树最多有 3000 个节点。
- 每一个节点值的范围是 [1, 1000] 
