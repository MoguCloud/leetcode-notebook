# [881. 救生艇](https://leetcode-cn.com/problems/boats-to-save-people/)
## Description
第 `i` 个人的体重为 `people[i]`，每艘船可以承载的最大重量为 `limit`。  
每艘船最多可同时载两人，但条件是这些人的重量之和最多为 `limit`。  
返回载到每一个人所需的最小船数。(保证每个人都能被船载)。  
## Example
### Example 1
Input:  
```
people = [1,2], limit = 3
```
Output:
```
1
```
1 艘船载 (1, 2)
### Example 2
Input:  
```
people = [3,2,2,1], limit = 3
```
Output:
```
3
```
3 艘船分别载 (1, 2), (2) 和 (3)
### Example 3
Input:  
```
people = [3,5,3,4], limit = 5
```
Output:
```
4
```
4 艘船分别载 (3), (3), (4), (5)
## Hint
- 1 <= people.length <= 50000
- 1 <= people[i] <= limit <= 30000
