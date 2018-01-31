package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Len() != 3 {
		t.Errorf("栈插入异常， 长度应该为3")
	}

	if s.Peek() != 3 || s.Pop() != 3 ||
		s.Peek() != 2 || s.Pop() != 2 ||
		s.Peek() != 1 || s.Pop() != 1 || s.Len() != 0 {
		t.Errorf("栈弹出异常, 没有按插入顺序返回")
	}

	s.Push(1)
	if s.Pop() != 1 {
		t.Errorf("栈弹出异常, 没有按插入顺序返回")
	}
}
