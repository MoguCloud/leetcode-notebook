# [1452. 收藏清单](https://leetcode-cn.com/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list/)
## Description
给你一个数组 `favoriteCompanies` ，其中 `favoriteCompanies[i]` 是第 `i` 名用户收藏的公司清单**（下标从 0 开始）**。  
请找出不是其他任何人收藏的公司清单的子集的收藏清单，并返回该清单下标。下标需要按升序排列。
## Example
### Example 1
Input:  
```
favoriteCompanies = [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]
```
Output:
```
[0,1,4]
```
favoriteCompanies[2]=["google","facebook"] 是 favoriteCompanies[0]=["leetcode","google","facebook"] 的子集。  
favoriteCompanies[3]=["google"] 是 favoriteCompanies[0]=["leetcode","google","facebook"] 和 favoriteCompanies[1]=["google","microsoft"] 的子集。  
其余的收藏清单均不是其他任何人收藏的公司清单的子集，因此，答案为 [0,1,4] 。
### Example 2
Input:  
```
favoriteCompanies = [["leetcode","google","facebook"],["leetcode","amazon"],["facebook","google"]]
```
Output:
```
[0,1]
```
favoriteCompanies[2]=["facebook","google"] 是 favoriteCompanies[0]=["leetcode","google","facebook"] 的子集，因此，答案为 [0,1] 。
### Example 3
Input:  
```
favoriteCompanies = [["leetcode"],["google"],["facebook"],["amazon"]]
```
Output:
```
[0,1,2,3]
```
## Hint
- 1 <= favoriteCompanies.length <= 100
- 1 <= favoriteCompanies[i].length <= 500
- 1 <= favoriteCompanies[i][j].length <= 20
- favoriteCompanies[i] 中的所有字符串 **各不相同** 。
- 用户收藏的公司清单也 **各不相同** ，也就是说，即便我们按字母顺序排序每个清单， `favoriteCompanies[i] != favoriteCompanies[j]` 仍然成立。
- 所有字符串仅包含小写英文字母。
