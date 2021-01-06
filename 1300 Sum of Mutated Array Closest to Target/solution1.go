// Your runtime beats 81.82 % of golang submissions (12 ms)
// Your memory usage beats 100 % of golang submissions (5.3 MB)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	l := len(arr)
	m := target / l
	// 判断返回值是否小于arr[0]
	if m <= arr[0] {
		if abs(target-m*l) <= abs(target-(m+1)*l) {
			return m
		} else {
			return m + 1
		}
	}
	sum := arr[0]
	for i := 1; i < len(arr); i++ {
		if sum+arr[i]*(l-i) >= target {
			//返回值在(arr[i-1], arr[i]]中
			if sum+arr[i]*(l-i) == target {
				return arr[i]
			} else {
				// 二分查找
				s := arr[i-1]
				b := arr[i]
				m := (s + b) / 2
				for !(sum+m*(l-i) <= target && sum+(m+1)*(l-i) >= target) {
					if sum+m*(l-i) < target {
						s = m
					} else {
						b = m
					}
					m = (s + b) / 2
				}
				if abs(target-m*(l-i)-sum) <= abs(target-(m+1)*(l-i)-sum) {
					return m
				} else {
					return m + 1
				}
			}
		}
		sum += arr[i]
	}
	return arr[l-1]
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
