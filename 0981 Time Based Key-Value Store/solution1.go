// Your runtime beats 27.10 % of golang submissions (468 ms)
// Your memory usage beats 76.63 % of golang submissions (67.6 MB)
//
// 哈希表 + 二分搜索
// 哈希表 键为key 值为时间戳和值的二元组数组
// 查找时候使用二分搜索
//
// 时间复杂度 get O(logn) set O(1)
// 空间复杂度 O(n)

type Value struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]Value
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]Value),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	arr, ok := this.m[key]
	if !ok {
		arr = []Value{}
	}
	v := Value{
		timestamp: timestamp,
		value:     value,
	}
	arr = append(arr, v)
	this.m[key] = arr
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr, ok := this.m[key]
	if !ok {
		return ""
	} else {
		if arr[0].timestamp > timestamp {
			return ""
		}
		if arr[len(arr)-1].timestamp <= timestamp {
			return arr[len(arr)-1].value
		}
		i := 0
		j := len(arr) - 1
		for i < j {
			m := (i + j) / 2
			if arr[m].timestamp <= timestamp && arr[m+1].timestamp > timestamp {
				return arr[m].value
			} else if arr[m].timestamp > timestamp {
				j = m - 1
			} else {
				i = m + 1
			}
		}
		return ""
	}
}
