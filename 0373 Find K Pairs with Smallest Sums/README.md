# [373. 查找和最小的 K 对数字](https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/)
## Description
给定两个以 **升序排列** 的整数数组 `nums1` 和 `nums2` , 以及一个整数 `k` 。  
定义一对值 `(u,v)`，其中第一个元素来自 `nums1`，第二个元素来自 `nums2` 。  
请找到和最小的 `k` 个数对 `(u1,v1)`,  `(u2,v2)`  ...  `(uk,vk)` 。
## Example
### Example 1
Input:  
```
nums1 = [1,7,11], nums2 = [2,4,6], k = 3
```
Output:
```
[1,2],[1,4],[1,6]
```
返回序列中的前 3 对数：  
[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
### Example 2
Input:  
```
nums1 = [1,1,2], nums2 = [1,2,3], k = 2
```
Output:
```
[1,1],[1,1]
```
返回序列中的前 2 对数：  
[1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
### Example 3
Input:  
```
nums1 = [1,2], nums2 = [3], k = 3 
```
Output:
```
[1,3],[2,3]
```
也可能序列中所有的数对都被返回:[1,3],[2,3]
## Hint
- 1 <= nums1.length, nums2.length <= 10^5
- -10^9 <= nums1[i], nums2[i] <= 10^9
- nums1 和 nums2 均为升序排列
- 1 <= k <= 1000
