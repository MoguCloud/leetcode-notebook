// Your runtime beats 64.29 % of golang submissions (16 ms)
// Your memory usage beats 51.79 % of golang submissions (12.9 MB)
//
// DFS
// 将每个子树通过字符串序列化
// 再使用哈希表 键为 每个子树序列化之后的字符串， 值为出现的次数
// 如果出现次数=2次，则是重复的子树

var ret []*TreeNode
var memo map[string]int

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ret = []*TreeNode{}
	memo = make(map[string]int)
	dfs(root)
	return ret
}

func dfs(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	encodedStr := fmt.Sprintf("%v,%v,%v", root.Val, dfs(root.Left), dfs(root.Right))
	memo[encodedStr] += 1
	if count, _ := memo[encodedStr]; count == 2 {
		ret = append(ret, root)
	}
	return encodedStr
}
