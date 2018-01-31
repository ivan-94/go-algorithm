// Package set使用链表的形式来实现集合方法
package set

import (
	"github.com/carney520/go-algorithm/data-structure/list"
)

var (
	ErrConflict = errors.New("元素已存在")
)

type Comparator = func(a, b interface{}) bool
type Set struct {
	list *list.List
	comparator Comparator
}


func (s *Set) find(v interface{}) (ret interface{}) {
	if s.list.Len() == 0 {
		return
	}
	
	s.list.Each(func(val interface, index int) bool {
		if s.comparator(v, val) {
			ret = val
			return true
		}
		return false
	})
	return
}

func (s *Set) Insert(data interface) error {
	if s.find(data) == nil {
		s.list.Append(data)
		return nil
	}
	return
}

func DefaultMatch(a, b interface{}) bool {
	return a == b
}

func New(match Comparator) *Set{
	return &Set{
		list: list.New()
	}
}
