// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 31.92 % of golang submissions (2.6 MB)
//
// 滑动窗口
//
// 时间复杂度 O(n)
// 空间复杂度 O(1)

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	// byteCount 用于存放各字母在s2窗口比s1多的个数
	// 如果'a'在s2窗口出现了3次，在s1中出现了1次，则byteCount[0]=2
	var byteCount [26]int
	for i := 0; i < len(s1); i++ {
		byteCount[s1[i]-'a']--
		byteCount[s2[i]-'a']++
	}
	// diff 用于记录s2窗口和s1数量不同的元素种类
	// 如果'a'在s2窗口出现了2次，在s1中出现了1次
	//     'b'在s2窗口出现了0次，在s1中出现了2次
	//     'c'在s2窗口出现了1次，在s1中出现了0次
	// 则diff=3
	diff := 0
	for i := 0; i < len(byteCount); i++ {
		if byteCount[i] != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		if s2[i] == s2[i-len(s1)] {
			continue
		}
		// 左指针向右移动一位时
		// 如果左指针指向的元素在s2和s1中数量相同，则不同元素种类diff会变大1
		// 如果左指针指向的元素在s2比s1中数量多1，则不同元素种类diff会变小1
		if byteCount[s2[i-len(s1)]-'a'] == 0 {
			diff++
		} else if byteCount[s2[i-len(s1)]-'a'] == 1 {
			diff--
		}
		byteCount[s2[i-len(s1)]-'a']--
		// 右指针向右移动一位时
		// 如果右指针指向的元素在s2和s1中数量相同，则不同元素种类diff会变大1
		// 如果右指针指向的元素在s2和s1中数量少1，则不同元素种类diff会变小1
		if byteCount[s2[i]-'a'] == -1 {
			diff--
		} else if byteCount[s2[i]-'a'] == 0 {
			diff++
		}
		byteCount[s2[i]-'a']++
		if diff == 0 {
			return true
		}
	}
	return false
}
