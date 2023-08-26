// Your runtime beats 16.67 % of golang submissions (104 ms)
// Your memory usage beats 11.11 % of golang submissions (16.6 MB)
//
// 拓扑排序
// 先根据 group 拓扑排序
// 然后每个 group 中的 item 在进行拓扑排序
//
// 时间复杂度 O(m+n)
// 空间复杂度 O(m+n)

type Group struct {
	id       int
	items    []int
	depends  map[int]struct{} // 依赖项 key groupID
	indegree int              // 入度 依赖项的数量
	after    map[int]struct{} // 被依赖项 key groupID
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	indegree := make([]int, n) // 每个item在当前group的入度
	after := make([][]int, n)  // 被依赖项
	// getGroupID 根据 item 获取 group_id，给 group 是 -1 的分配一个单独的 group_id
	getGroupID := func(item int) int {
		if group[item] != -1 {
			return group[item]
		}
		return -1 - item
	}
	// 根据 id 记录 *Group
	groupMap := make(map[int]*Group)
	// 初始化 groupMap
	for i := 0; i < n; i++ {
		groupID := getGroupID(i)
		g, ok := groupMap[groupID]
		if !ok {
			g = &Group{
				id:      groupID,
				items:   []int{},
				depends: make(map[int]struct{}),
				after:   make(map[int]struct{}),
			}
		}
		groupMap[groupID] = g
	}
	// 记录 group 的依赖项，item 和被依赖项以及入度等信息
	for i := 0; i < n; i++ {
		groupID := getGroupID(i)
		g := groupMap[groupID]
		g.items = append(g.items, i)
		for _, item := range beforeItems[i] {
			// 计算每个item在当前group的入度
			if getGroupID(item) == getGroupID(i) {
				after[item] = append(after[item], i)
				indegree[i]++
			}
			// 依赖项是在当前group的话则不记录
			if groupID == getGroupID(item) {
				continue
			}
			g.depends[getGroupID(item)] = struct{}{}
			itemGroup := groupMap[getGroupID(item)]
			itemGroup.after[groupID] = struct{}{}
		}
	}
	// 根据group拓扑排序
	groupOrder := []*Group{}
	for _, g := range groupMap {
		if len(g.depends) == 0 {
			groupOrder = append(groupOrder, g)
		}
		g.indegree = len(g.depends)
	}
	for i := 0; i < len(groupOrder); i++ {
		for id, _ := range groupOrder[i].after {
			afterGroup := groupMap[id]
			afterGroup.indegree--
			if afterGroup.indegree == 0 {
				groupOrder = append(groupOrder, afterGroup)
			}
		}
	}
	ret := []int{}
	for _, group := range groupOrder {
		// group内根据item的依赖进行拓扑排序
		itemOrder := []int{}
		for _, item := range group.items {
			if indegree[item] == 0 {
				itemOrder = append(itemOrder, item)
			}
		}
		for i := 0; i < len(itemOrder); i++ {
			for _, item := range after[itemOrder[i]] {
				if getGroupID(item) == getGroupID(itemOrder[i]) {
					indegree[item]--
				}
				if indegree[item] == 0 {
					itemOrder = append(itemOrder, item)
				}
			}
		}
		ret = append(ret, itemOrder...)
	}
	// 如果最后排序完的结果数量与 item 数量不一样，则没有解决方案
	if len(ret) != n {
		return []int{}
	}
	return ret
}
