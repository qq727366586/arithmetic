package __1

import "errors"

var (
	errEmptyStack = errors.New("stack is empty")
)

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

// 入栈操作
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 出栈操作
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyStack
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item, nil
}

// 返回栈顶元素
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errEmptyStack
	}

	return s.data[len(s.data)-1], nil
}

// 返回栈的大小
func (s *Stack) Size() int {
	return len(s.data)
}
