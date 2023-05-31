// Your runtime beats 50 % of golang submissions (63 ms)
// Your memory usage beats 51.89 % of golang submissions (8 MB)
//
// 并查集

type UnionFind struct {
	parent     []int
	connection int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{parent: make([]int, n), connection: n}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Union(i, j int) {
	if uf.Find(i) != uf.Find(j) {
		uf.parent[uf.Find(i)] = uf.Find(uf.parent[j])
		uf.connection--
	}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func accountsMerge(accounts [][]string) [][]string {
	mailIdxMap := make(map[string]int)
	idxNameMap := make(map[int]string)
	idx := 0
	for _, account := range accounts {
		name := account[0]
		for i := 1; i < len(account); i++ {
			idxNameMap[idx] = name
			_, ok := mailIdxMap[account[i]]
			if !ok {
				mailIdxMap[account[i]] = idx
				idx++
			}
		}
	}
	uf := NewUnionFind(idx)
	for _, account := range accounts {
		if len(account) > 2 {
			idx1 := mailIdxMap[account[1]]
			for i := 2; i < len(account); i++ {
				idx2 := mailIdxMap[account[i]]
				uf.Union(idx1, idx2)
			}
		}
	}
	unionIdx := 0
	unionIdxMap := make(map[int]int)
	ret := make([][]string, uf.connection)
	for mail, idx := range mailIdxMap {
		parent := uf.Find(idx)
		uIdx, ok := unionIdxMap[parent]
		if !ok {
			uIdx = unionIdx
			unionIdx++
			unionIdxMap[parent] = uIdx
			ret[uIdx] = []string{idxNameMap[idx]}
		}
		ret[uIdx] = append(ret[uIdx], mail)
	}
	for _, r := range ret {
		sort.Strings(r[1:])
	}
	return ret
}
