# [735. 行星碰撞](https://leetcode.cn/problems/asteroid-collision/)
## Description
给定一个整数数组 `asteroids`，表示在同一行的行星。  
对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。  
找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。
## Example
### Example 1
Input:  
```
asteroids = [5,10,-5]
```
Output:
```
[5,10]
```
10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
### Example 2
Input:  
```
asteroids = [8,-8]
```
Output:
```
[]
```
8 和 -8 碰撞后，两者都发生爆炸。
### Example 3
Input:  
```
asteroids = [10,2,-5]
```
Output:
```
[10]
```
2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。
## Hint
- 2 <= asteroids.length <= 10^4
- -1000 <= asteroids[i] <= 1000
- asteroids[i] != 0
