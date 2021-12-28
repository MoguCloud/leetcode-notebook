# [875. 爱吃香蕉的珂珂](https://leetcode-cn.com/problems/koko-eating-bananas/)
## Description
珂珂喜欢吃香蕉。这里有 `N` 堆香蕉，第 `i` 堆中有 `piles[i]` 根香蕉。警卫已经离开了，将在 `H` 小时后回来。  
珂珂可以决定她吃香蕉的速度 `K` （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 `K` 根。如果这堆香蕉少于 `K` 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。    
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。  
返回她可以在 `H` 小时内吃掉所有香蕉的最小速度 `K`（`K` 为整数）。
## Example
### Example 1
Input:  
```
piles = [3,6,7,11], H = 8
```
Output:
```
4
```
### Example 2
Input:  
```
piles = [30,11,23,4,20], H = 5
```
Output:
```
30
```
### Example 3
Input:  
```
piles = [30,11,23,4,20], H = 6
```
Output:
```
23
```
## Hint
- 1 <= piles.length <= 10^4
- piles.length <= H <= 10^9
- 1 <= piles[i] <= 10^9