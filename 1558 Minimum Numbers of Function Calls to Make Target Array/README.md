# [1558. 得到目标数组的最少函数调用次数](https://leetcode-cn.com/problems/minimum-numbers-of-function-calls-to-make-target-array/)
## Description
![](https://assets.leetcode.com/uploads/2020/07/10/sample_2_1887.png)  
给你一个与 `nums` 大小相同且初始值全为 `0` 的数组 `arr` ，请你调用以上函数得到整数数组 `nums` 。  
请你返回将 `arr` 变成 `nums` 的最少函数调用次数。  
答案保证在 `32` 位有符号整数以内。
## Example
### Example 1
Input:  
```
nums = [1,5]
```
Output:
```
5
```
给第二个数加 1 ：[0, 0] 变成 [0, 1] （1 次操作）。  
将所有数字乘以 2 ：[0, 1] -> [0, 2] -> [0, 4] （2 次操作）。  
给两个数字都加 1 ：[0, 4] -> [1, 4] -> [1, 5] （2 次操作）。  
总操作次数为：1 + 2 + 2 = 5 。
### Example 2
Input:  
```
nums = [2,2]
```
Output:
```
3
```
给两个数字都加 1 ：[0, 0] -> [0, 1] -> [1, 1] （2 次操作）。  
将所有数字乘以 2 ： [1, 1] -> [2, 2] （1 次操作）。  
总操作次数为： 2 + 1 = 3 。  
### Example 3
Input:
```
nums = [4,2,5]
```
Output:
```
6
```
（初始）[0,0,0] -> [1,0,0] -> [1,0,1] -> [2,0,2] -> [2,1,2] -> [4,2,4] -> [4,2,5] （nums 数组）。
### Example 4
Input:
```
nums = [3,2,2,4]
```
Output:
```
7
```
### Example 5
Input:
```
nums = [2,4,8,16]
```
Output:
```
8
```
## Hint
- 1 <= nums.length <= 10^5
- 0 <= nums[i] <= 10^9

