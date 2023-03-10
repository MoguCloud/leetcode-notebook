// Your runtime beats 90.91 % of golang submissions (12 ms)
// Your memory usage beats 96.97 % of golang submissions (7.2 MB)
//
// 蓄水池算法
// 当链表前进到第一个节点时，等于该值的概率为1/1，
// 当链表前进到第二个节点时，等于该值的概率为1/2，等于之前值概率为1 - 1/2 = 1/2
// 当链表前进到第三个节点时，等于该值的概率为1/3，等于之前值概率为1 - 1/3 = 2/3
// 当链表前进到第n个几点时，等于该值的概率为1/n，等于之前值概率为1 - 1/n = (n-1)/n
// 所以前进到第n个节点时，有1/n的概率将结果置为当前节点，当链表遍历完，结果就是需要的随机数。
//
// 时间复杂度 初始化 O(1) 随机选择 O(n)
// 空间复杂度 O(1)

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	i := 1
	ret := 0
	for head := this.head; head != nil; head = head.Next {
		if rand.Intn(i) == 0 {
			ret = head.Val
		}
		i += 1
	}
	return ret
}
