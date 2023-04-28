// Your runtime beats 62.50 % of golang submissions (31 ms)
// Your memory usage beats 18.75 % of golang submissions (7.7 MB)
//
// 并查集
//
// 时间复杂度 O(n*n*m+nlogn) n为字符串个数，m为字符串长度
// 空间复杂度 O(n)

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{parent: make([]int, n)}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Union(i, j int) {
	pi := uf.Find(i)
	pj := uf.Find(j)
	if pi != pj {
		uf.parent[pi] = pj
	}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := NewUnionFind(n)
	ret := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if uf.Find(i) != uf.Find(j) && isSimilar(strs[i], strs[j]) {
				uf.Union(i, j)
				ret--
			}
		}
	}
	return ret
}

func isSimilar(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}
	diff := []int{}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			diff = append(diff, i)
		}
		if len(diff) > 2 {
			return false
		}
	}
	if len(diff) == 2 && str1[diff[0]] == str2[diff[1]] && str2[diff[0]] == str1[diff[1]] {
		return true
	}
	return false
}
