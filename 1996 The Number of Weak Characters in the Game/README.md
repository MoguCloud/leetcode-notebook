# [1996. 游戏中弱角色的数量](https://leetcode-cn.com/problems/the-number-of-weak-characters-in-the-game/)
## Description
你正在参加一个多角色游戏，每个角色都有两个主要属性：**攻击** 和 **防御** 。给你一个二维整数数组 `properties` ，其中 `properties[i] = [attacki, defensei]` 表示游戏中第 `i` 个角色的属性。  
如果存在一个其他角色的攻击和防御等级 **都严格高于** 该角色的攻击和防御等级，则认为该角色为 **弱角色** 。更正式地，如果认为角色 `i` **弱于** 存在的另一个角色 `j` ，那么 `attack[j] > attack[i]` 且 `defense[j] > defense[i]` 。  
返回 弱角色 的数量。
## Example
### Example 1
Input:  
```
properties = [[5,5],[6,3],[3,6]]
```
Output:
```
0
```
不存在攻击和防御都严格高于其他角色的角色。
### Example 2
Input:  
```
properties = [[2,2],[3,3]]
```
Output:
```
1
```
第一个角色是弱角色，因为第二个角色的攻击和防御严格大于该角色。
### Example 3
Input:  
```
properties = [[1,5],[10,4],[4,3]]
```
Output:
```
1
```
第三个角色是弱角色，因为第二个角色的攻击和防御严格大于该角色。
## Hint
- 2 <= properties.length <= 10^5
- properties[i].length == 2
- 1 <= attacki, defensei <= 10^5
