# [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/description/)
## Description
给你一个整数数组 `nums`，有一个大小为 `k` 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 `k` 个数字。滑动窗口每次只向右移动一位。  
返回 *滑动窗口中的最大值* 。  
## Example
### Example 1
Input:  
```
nums = [1,3,-1,-3,5,3,6,7], k = 3
```
Output:
```
[3,3,5,5,6,7]
```
滑动窗口的位置                最大值  
---------------               -----  
[1  3  -1] -3  5  3  6  7       3  
 1 [3  -1  -3] 5  3  6  7       3  
 1  3 [-1  -3  5] 3  6  7       5  
 1  3  -1 [-3  5  3] 6  7       5  
 1  3  -1  -3 [5  3  6] 7       6  
 1  3  -1  -3  5 [3  6  7]      7  
### Example 2
Input:  
```
nums = [1], k = 1
```
Output:
```
[1]
```
## Hint
- 1 <= nums.length <= 10^5
- -10^4 <= nums[i] <= 10^4
- 1 <= k <= nums.length
