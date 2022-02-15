// Your runtime beats 14.29 % of golang submissions (443 ms)
// Your memory usage beats 85.71 % of golang submissions (22.7 MB)
//
// 使用并查集将 allowedSwaps 中连通的 pair 合并成连通分支
// 再针对不同的连通分支进行距离计算
// 对于每个连通分支，使用2个map分别记录source和target中每个元素出现的次数
// 次数差为该连通分支的距离
// 所有连通分支的距离和即为答案

type UnionFind struct {
	arr []int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		arr: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.arr[i] = i
	}
	return u
}

func (u *UnionFind) find(x int) int {
	if u.arr[x] != x {
		u.arr[x] = u.find(u.arr[x])
	}
	return u.arr[x]
}

func (u *UnionFind) union(i, j int) {
	pi := u.find(i)
	pj := u.find(j)
	if pi != pj {
		u.arr[pi] = pj
	}
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	u := NewUnionFind(len(source))
	for _, pair := range allowedSwaps {
		u.union(pair[0], pair[1])
	}
	m := make(map[int][]int) // key 并查集父id value 并查集父id对应的连通分支
	for i := 0; i < len(source); i++ {
		parent := u.find(i)
		arr, ok := m[parent]
		if !ok {
			arr = []int{}
		}
		arr = append(arr, i)
		m[parent] = arr
	}
	// 每个连通分支计算距离
	ret := 0
	for _, arr := range m {
		countSource := make(map[int]int)
		countTarget := make(map[int]int)
		// 计算source和target中每个元素出现次数
		for _, i := range arr {
			cs, _ := countSource[source[i]]
			cs += 1
			countSource[source[i]] = cs
			ct, _ := countTarget[target[i]]
			ct += 1
			countTarget[target[i]] = ct
		}
		// 计算元素出现次数的差，即连通分支的距离
		c := len(arr)
		for k, cs := range countSource {
			ct, _ := countTarget[k]
			c -= min(cs, ct)
		}
		ret += c
	}
	return ret
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
