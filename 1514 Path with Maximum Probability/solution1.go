// Your runtime beats 6.38 % of golang submissions (230 ms)
// Your memory usage beats 6.38 % of golang submissions (13.4 MB)
//
// Dijkstra
//
// 时间复杂度 O(mlogm) m 为边的数量
// 空间复杂度 O(m)

type pair struct {
	id   int
	prob float64
}

type pairHeap []pair

func (h pairHeap) Len() int {
	return len(h)
}

func (h pairHeap) Less(i, j int) bool {
	return h[i].prob > h[j].prob
}

func (h pairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pairHeap) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *pairHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	neighbor := make([][]int, n)
	probMap := make(map[string]float64)
	for i := 0; i < len(edges); i++ {
		neighbor[edges[i][0]] = append(neighbor[edges[i][0]], edges[i][1])
		neighbor[edges[i][1]] = append(neighbor[edges[i][1]], edges[i][0])
		probMap[fmt.Sprintf("%d %d", edges[i][0], edges[i][1])] = succProb[i]
		probMap[fmt.Sprintf("%d %d", edges[i][1], edges[i][0])] = succProb[i]
	}
	dis := make([]float64, n)
	dis[start] = 1
	h := &pairHeap{pair{start, 1}}
	heap.Init(h)
	for h.Len() > 0 {
		cur := heap.Pop(h).(pair)
		if cur.id == end {
			return cur.prob
		}
		if cur.prob < dis[cur.id] {
			continue
		}
		for _, nei := range neighbor[cur.id] {
			toProb := probMap[fmt.Sprintf("%d %d", nei, cur.id)]
			if toProb == 0 {
				continue
			}
			neiProb := cur.prob * toProb
			if neiProb > dis[nei] {
				dis[nei] = neiProb
				heap.Push(h, pair{nei, neiProb})
			}
		}
	}
	return 0
}
