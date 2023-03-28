// Your runtime beats 94.12 % of golang submissions (62 ms)
// Your memory usage beats 94.12 % of golang submissions (10.2 MB)
//
// 并查集
//
// 时间复杂度 O(n+m×α(m))，其中 nn 是计算机数，mm 是线缆数，αα 是反阿克曼函数
// 空间复杂度 O(n)

type UF struct {
	parent []int
	count  int
}

func NewUF(n int) *UF {
	uf := &UF{
		parent: make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UF) Union(i int, j int) {
	if uf.Find(i) == uf.Find(j) {
		return
	}
	uf.parent[uf.Find(i)] = uf.Find(j)
	uf.count--
}

func (uf *UF) Find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.Find(uf.parent[i])
	return uf.parent[i]
}

func makeConnected(n int, connections [][]int) int {
	// n 台计算机需要至少 n - 1 条线缆才能连通
	if len(connections) < n-1 {
		return -1
	}
	// 使用并查集计算连通分量
	uf := NewUF(n)
	for _, c := range connections {
		uf.Union(c[0], c[1])
	}
	if uf.count == 1 {
		return 0
	}
	// 连通分量为1时，所有计算机连通，所以需要插拔连通分量 - 1 条线缆，使连通分量为1
	return uf.count - 1

}
