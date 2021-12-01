# [894. 所有可能的满二叉树](https://leetcode-cn.com/problems/all-possible-full-binary-trees/)
## Description
**满二叉树** 是一类二叉树，其中每个结点恰好有 `0` 或 `2` 个子结点。  
返回包含 `N` 个结点的所有可能满二叉树的列表。 答案的每个元素都是一个可能树的根结点。  
答案中每个树的每个 **结点** 都 **必须** 有 `node.val=0`。  
你可以按任何顺序返回树的最终列表。  
## Example
### Example 1
Input:  
```
7
```
Output:
```
[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
```
![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/22/fivetrees.png)  
## Hint
- 1 <= n <= 20
