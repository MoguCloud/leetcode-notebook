// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 24.75 % of golang submissions (2.6 MB)
// 
// 先排序
// 计算num[0:0] 的字串 再计算num[0:1]的字串 一直到 num[0:len(num)]的字串。
// 通过一个哈希表记录已存在的字串

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{[]int{}}
	used := make(map[string]struct{})
	for i := 0; i < len(nums); i++ {
		length := len(ret)
		for j := 0; j < length; j++ {
			buf := []byte{}
			for k := 0; k < len(ret[j]); k++ {
				buf = append(buf, byte(ret[j][k]))
				buf = append(buf, '#')
			}
			buf = append(buf, byte(nums[i]))
			str := string(buf)
			if _, ok := used[str]; !ok {
				arr := make([]int, len(ret[j])+1)
				copy(arr, ret[j])
				arr[len(arr)-1] = nums[i]
				ret = append(ret, arr)
				used[str] = struct{}{}
			}
		}
	}
	return ret
}
