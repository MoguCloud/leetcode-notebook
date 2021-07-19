// Your runtime beats 100 % of golang submissions (0 ms)
// Your memory usage beats 66.67 % of golang submissions (2 MB)
//
// Stack
//
// 定义一个栈
// 遍历字符串，当遇到字符不为']'时，不断压入栈中。
// 当遇到字符为']'时，不断从栈中弹出栈顶字符，并存入一个byte数组
// 当栈顶元素为'['时，则说明从栈中弹出并存入byte数组的字符构成encoded_string
// 再从栈顶弹出，直到栈顶元素不为数字，此时弹出的byte构成重复次数k
// 将encoded_string重复k次后压入栈中，继续遍历字符串
//
// 时间复杂度 O(n)
// 空间复杂度 O(n)

type Stack struct {
	arr []byte
}

func NewStack() *Stack {
	return &Stack{arr: []byte{}}
}

func (s *Stack) Top() byte {
	return s.arr[len(s.arr)-1]
}

func (s *Stack) Push(b byte) {
	s.arr = append(s.arr, b)
}

func (s *Stack) Pop() byte {
	b := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return b
}

func (s *Stack) Len() int {
	return len(s.arr)
}

func decodeString(s string) string {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack.Push(s[i])
		} else {
			encoded_byte := []byte{}
			for {
				if stack.Top() == '[' {
					stack.Pop()
					break
				} else {
					encoded_byte = append(encoded_byte, stack.Pop())
				}
			}
			repeat_byte := []byte{}
			for {
				if stack.Len() == 0 || stack.Top() < '0' || stack.Top() > '9' {
					break
				} else {
					repeat_byte = append(repeat_byte, stack.Pop())
				}
			}
			for ii, jj := 0, len(repeat_byte)-1; ii < jj; ii, jj = ii+1, jj-1 {
				repeat_byte[ii], repeat_byte[jj] = repeat_byte[jj], repeat_byte[ii]
			}
			repeat, _ := strconv.Atoi(string(repeat_byte))
			for j := 0; j < repeat; j++ {
				for k := 1; k <= len(encoded_byte); k++ {
					stack.Push(encoded_byte[len(encoded_byte)-k])
				}
			}
		}
	}
	return string(stack.arr)
}
