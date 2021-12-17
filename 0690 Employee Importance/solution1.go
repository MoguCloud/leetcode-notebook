// Your runtime beats 96.47 % of golang submissions (7 ms)
// Your memory usage beats 9.41 % of golang submissions (7.4 MB)
//
// DFS
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

var employeeMap map[int]*Employee

func getImportance(employees []*Employee, id int) int {
	employeeMap = make(map[int]*Employee) // key id
	for _, employee := range employees {
		employeeMap[(*employee).Id] = employee
	}
	return getTotalImportance(employeeMap[id])
}

func getTotalImportance(employee *Employee) int {
	sum := (*employee).Importance
	for _, subordinateId := range (*employee).Subordinates {
		subordinate := employeeMap[subordinateId]
		sum += getTotalImportance(subordinate)
	}
	return sum
}
