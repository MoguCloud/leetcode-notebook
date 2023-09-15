// Your runtime beats 78.23 % of golang submissions (266 ms)
// Your memory usage beats 90.48 % of golang submissions (21.53 MB)
//
// Kruskal 算法
// 按费用从小到大排序各条边
// 遍历各条边，2个点没有连通则进行连通
// 直到只剩一个连通分量
//
// 时间复杂度 O(n^2logn)
// 空间复杂度 O(n^2)

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func (uf *UnionFind) connected(i, j int) bool {
	rootI := uf.find(i)
	rootJ := uf.find(j)
	return rootI == rootJ
}

func (uf *UnionFind) union(i, j int) {
	rootI := uf.find(i)
	rootJ := uf.find(j)
	if rootI == rootJ {
		return
	}
	if uf.rank[i] < uf.rank[j] {
		uf.parent[rootI] = rootJ
		uf.rank[rootI] += uf.rank[rootJ]
	} else {
		uf.parent[rootJ] = rootI
		uf.rank[rootJ] += uf.rank[rootI]
	}
	uf.count--
}

func (uf *UnionFind) find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.find(uf.parent[i])
	return uf.parent[i]
}

func minCostConnectPoints(points [][]int) int {
	distances := make([][3]int, len(points)*(len(points)-1)/2)
	k := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pi := points[i]
			pj := points[j]
			dis := abs(pi[0]-pj[0]) + abs(pi[1]-pj[1])
			distances[k] = [3]int{i, j, dis}
			k++
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i][2] < distances[j][2]
	})
	k = 0
	uf := NewUnionFind(len(points))
	ret := 0
	for uf.count != 1 {
		if !uf.connected(distances[k][0], distances[k][1]) {
			uf.union(distances[k][0], distances[k][1])
			ret += distances[k][2]
		}
		k++
	}
	return ret
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
