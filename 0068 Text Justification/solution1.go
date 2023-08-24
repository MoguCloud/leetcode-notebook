// Your runtime beats 73.88 % of golang submissions (1 ms)
// Your memory usage beats 70.78 % of golang submissions (2.16 MB)
//
// 模拟
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

func fullJustify(words []string, maxWidth int) []string {
	ret := []string{}
	i, j := 0, 0
	for i < len(words) {
		// 因为计算长度时需要计算2个单词中间的空格，所以width=-1
		width := -1
		j = i
		// 用双指针找出一组满足不大于最大长度的单词
		for j < len(words) {
			// 因为需要计算2个单词中间的空格，所以+1
			if width+len(words[j])+1 <= maxWidth {
				width += len(words[j]) + 1
				j++
			} else {
				break
			}
		}
		var row []string
		wordCount := j - i
		if j == len(words) {
			// 最后一行，每个单词中间一个空格，多余的空格添加在最末尾
			word := strings.Join(words[i:j], " ")
			row = make([]string, 2)
			row[0] = word
			row[1] = strings.Repeat(" ", maxWidth-len(word))
			ret = append(ret, strings.Join(row, ""))
		} else {
			if width < maxWidth {
				if wordCount == 1 {
					// 一行只有一个单词，需要右填充空格
					row = make([]string, 2)
					row[0] = words[i]
					row[1] = strings.Repeat(" ", maxWidth-len(words[i]))
					ret = append(ret, strings.Join(row, ""))
				} else {
					// 计算该行需要的空格数量
					spaces := maxWidth
					for k := i; k < j; k++ {
						spaces -= len(words[k])
					}
					// spaceCounts[k] 为 单词 words[i] 与words [i+k] 中间的空格
					spaceCounts := make([]int, wordCount-1)
					// row为该行的结果 [ word[i], spaceCounts[0]个空格, word[i+1], spaceCounts[1]个空格... spaceCounts[wrodCount-1]个空格， word[j - 1] ]
					row = make([]string, wordCount*2-1)
					// 先将空格数量平均分到每个spaceCounts[k]中，平均分完后多出的空格后续再进行分配
					for k := 0; k < len(spaceCounts); k++ {
						spaceCounts[k] = spaces / (wordCount - 1)
					}
					// 计算平均分配完空格后多余空格的数量
					spaces = spaces - spaces/(wordCount-1)*(wordCount-1)
					// 将该行的字母插入结果数组 row 中
					for k := 0; k*2 < len(row); k++ {
						row[k*2] = words[i+k]
					}
					for k := 0; k < len(spaceCounts); k++ {
						// 分配平均分配完多出的空格
						if spaces > 0 {
							spaceCounts[k]++
							spaces--
						}
						// 将空格插入结果数组中
						row[k*2+1] = strings.Repeat(" ", spaceCounts[k])
					}
					ret = append(ret, strings.Join(row, ""))
				}
			} else {
				// 正好填满一行
				ret = append(ret, strings.Join(words[i:j], " "))
			}
		}
		i = j
	}
	return ret
}
