# [662. 二叉树最大宽度](https://leetcode-cn.com/problems/maximum-width-of-binary-tree/)
## Description
给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与**满二叉树（full binary tree）**结构相同，但一些节点为空。  
每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
## Example
### Example 1
Input:  
```
           1
         /   \
        3     2
       / \     \  
      5   3     9 
```
Output:
```
4
```
最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
### Example 2
Input:  
```
          1
         /  
        3    
       / \       
      5   3  
```
Output:
```
2
```
最大值出现在树的第 3 层，宽度为 2 (5,3)。
### Example 3
Input:  
```
          1
         / \
        3   2 
       /        
      5     
```
Output:
```
2
```
最大值出现在树的第 2 层，宽度为 2 (3,2)。
### Example 4
Input:  
```
          1
         / \
        3   2
       /     \  
      5       9 
     /         \
    6           7
```
Output:
```
8	
```
最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。
## Hint
- 答案在32位有符号整数的表示范围内。
