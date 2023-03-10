// Your runtime beats 84.85 % of golang submissions (13 ms)
// Your memory usage beats 9.9 % of golang submissions (8.3 MB)
//
// 用数组记录数据再随机取
//
// 时间复杂度 初始化 O(n) 随机选择 O(1)
// 空间复杂度 O(n)

type Solution struct {
	arr []*ListNode
}

func Constructor(head *ListNode) Solution {
	s := Solution{arr: []*ListNode{}}
	for head != nil {
		s.arr = append(s.arr, head)
		head = head.Next
	}
	return s
}

func (this *Solution) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))].Val
}
