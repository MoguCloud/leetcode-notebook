// 排序
//
// Your runtime beats 100 % of golang submissions (376 ms)
// Your memory usage beats 64.29 % of golang submissions (19.8 MB)
//
// 按攻击力降序排序，攻击力相同的按防御力升序排序。
// 然后遍历数组，维护一个最大防御力maxDef
// 如果当前角色的防御力 < maxDef
// 则之前存在一个角色攻击力 >= 当前角色 且 防御力 > 当前角色
// 但因为排序规则，攻击力相同时，防御力低的在前
// 所以之前存在一个角色攻击力 > 当前角色 且 防御力 > 当前角色
// 即当前角色为弱者 
//
// 时间复杂度 O(nlogn)
// 空间复杂度 O(n)

package main

import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] > properties[j][0] {
			return true
		} else if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		} else {
			return false
		}

	}) // 按atk降序排序 atk相同的按def升序排序

	count := 0
	maxDef := 0
	for i := 0; i < len(properties); i++ {
		if properties[i][1] < maxDef {
			count += 1
		} else {
			maxDef = properties[i][1]
		}
	}
	return count
}
