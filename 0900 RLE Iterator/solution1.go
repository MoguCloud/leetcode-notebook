// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 16.67 % of golang submissions (3.6 MB)

type RLEIterator struct {
	arr      []int
	offset   int  // 当前指向的元素
	left     int  // 当前元素剩余个数
	finished bool // 是否迭代结束
}

func Constructor(encoding []int) RLEIterator {
	r := RLEIterator{
		arr:      encoding,
		offset:   0,
		left:     encoding[0],
		finished: false,
	}
	return r
}

func (this *RLEIterator) Next(n int) int {
	if this.finished {
		return -1
	}
	// 当 n > this.left 时，当前元素剩余个数不足，需要迭代下一个元素，直到元素个数足够
	for n > this.left {
		this.offset += 2
		if this.offset >= len(this.arr) {
			this.finished = true
			return -1
		}
		n -= this.left
		this.left = this.arr[this.offset]
	}
	this.left -= n
	return this.arr[this.offset+1]
}
