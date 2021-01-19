// Your runtime beats 90.91 % of golang submissions (148 ms)
// Your memory usage beats 100 % of golang submissions (8.1 MB)

func maxProfit(inventory []int, orders int) int {
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i] > inventory[j]
	})
	threshold := int64(math.Pow(10, 9) + 7)
	var total int64
	idx := 0   // 用于记录inventory的前idx为相同的值
	delta := 0 // inventory[idx] 与 inventoy[idx+1] 的差值
	for orders > 0 {
		for idx < len(inventory)-1 {
			if inventory[idx] == inventory[idx+1] {
				idx += 1
			} else {
				break
			}
		}
		if idx < len(inventory)-1 {
			delta = inventory[idx] - inventory[idx+1]
		} else {
			delta = inventory[idx]
		}

		if (idx+1)*delta < orders {
			total += sum(inventory[0], delta) * int64(idx+1)
			orders -= (idx + 1) * delta
			// for i := 0; i <= idx; i++ {
			//     inventory[i] = inventory[idx+1]
			// }
			// 因为 1 ~ idx-1不会被使用 所以可以简化
			inventory[0] = inventory[idx+1]
			inventory[idx] = inventory[idx+1]
		} else {
			delta = orders / (idx + 1)
			mod := orders % (idx + 1)
			total += sum(inventory[0], delta) * int64(idx+1)
			inventory[0] -= delta
			total += int64(inventory[0] * mod)
			break
		}
	}
	if total > threshold {
		total = total % threshold
	}
	return int(total)
}

func sum(num int, i int) int64 {
	return (int64(num) + int64(num) - int64(i) + 1) * int64(i) / 2
}
