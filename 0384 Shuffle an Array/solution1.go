// Your runtime beats 93.10 % of golang submissions (44 ms)
// Your memory usage beats 63.22 % of golang submissions (9.3 MB)
//
// 从第0张到第len-1张牌开始遍历
// 第i次遍历的时候，选择从前i张牌(包含第i张牌)中选择一张，与第i张牌进行交换

type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
	return Solution{
		arr: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	copyArr := make([]int, len(this.arr))
	copy(copyArr, this.arr)
	for i := 1; i < len(this.arr); i++ {
		roll := rand.Intn(i + 1)
		copyArr[i], copyArr[roll] = copyArr[roll], copyArr[i]
	}
	return copyArr
}
