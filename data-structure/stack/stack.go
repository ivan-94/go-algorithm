package stack

import (
	"github.com/carney520/go-algorithm/data-structure/list"
)

// Stack 表示堆栈
type Stack struct {
	list *list.List
}

// Push 推进堆栈
func (s *Stack) Push(data interface{}) {
	s.list.Append(data)
}

// Len 获取栈的长度
func (s *Stack) Len() int {
	return s.list.Len()
}

// Peek 获取栈顶元素
func (s *Stack) Peek() interface{} {
	tail := s.list.Tail()
	if tail != nil {
		return tail.Data
	}
	return nil
}

// Pop 弹出栈顶元素
func (s *Stack) Pop() interface{} {
	tail := s.list.Tail()
	if tail != nil {
		s.list.Remove(tail)
		return tail.Data
	}
	return nil
}

// New 创建一个栈
func New() *Stack {
	return &Stack{
		list: list.New(),
	}
}
