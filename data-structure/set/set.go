// Package set 使用链表的形式来实现集合方法
package set

import (
	"errors"
	"fmt"
	"strings"

	"github.com/carney520/go-algorithm/data-structure/list"
)

// 错误码
var (
	ErrConflict = errors.New("元素已存在")
	ErrNotFound = errors.New("元素不存在")
)

// Comparator 比较函数类型
type Comparator = func(a, b interface{}) bool

// Set 表示集合类型
type Set struct {
	list       *list.List
	comparator Comparator
}

func (s *Set) find(v interface{}) (ret *list.Node) {
	if s.list.Len() == 0 {
		return
	}

	s.list.Each(func(node *list.Node, index int) bool {
		if s.comparator(v, node.Data) {
			ret = node
			return true
		}
		return false
	})
	return
}

// Insert 插入一个新成员
func (s *Set) Insert(data interface{}) error {
	if s.find(data) == nil {
		s.list.Append(data)
		return nil
	}
	return ErrConflict
}

// Has 判断成员是否存在
func (s *Set) Has(data interface{}) bool {
	return s.find(data) != nil
}

// Remove 移除成员
func (s *Set) Remove(data interface{}) error {
	n := s.find(data)
	if n != nil {
		// remove
		s.list.Remove(n)
		return nil
	}
	return ErrNotFound
}

// Len 返回集合的成员数
func (s *Set) Len() int {
	return s.list.Len()
}

// Clone 克隆一个新集合
func (s *Set) Clone() *Set {
	ns := New(s.comparator)
	if s.Len() == 0 {
		return ns
	}
	s.list.Each(func(node *list.Node, index int) bool {
		ns.list.Append(node.Data)
		return false
	})
	return ns
}

// Equal 比较两个集合是否相等
func (s *Set) Equal(v *Set) bool {
	if s.Len() != v.Len() {
		return false
	}

	if s.Len() == 0 {
		return true
	}

	eq := true
	s.list.Each(func(node *list.Node, index int) bool {
		if !v.Has(node.Data) {
			eq = false
			return true
		}
		return false
	})
	return eq
}

// Each 迭代集合
func (s *Set) Each(it func(data interface{}, i int) (stop bool)) {
	i := 0
	s.list.Each(func(n *list.Node, _ int) bool {
		rt := it(n.Data, i)
		i++
		return rt
	})
}

// Union 并集
func (s *Set) Union(v *Set) *Set {
	if s == v {
		return s
	}
	rt := s.Clone()
	v.Each(func(data interface{}, i int) bool {
		rt.Insert(data)
		return false
	})
	return rt
}

// Intersection 求交集
func (s *Set) Intersection(v *Set) *Set {
	n := New(s.comparator)
	if s.Len() == 0 || v.Len() == 0 {
		return n
	}
	mins, maxs := s, v
	if s.Len() > v.Len() {
		mins, maxs = v, s
	}

	mins.Each(func(data interface{}, i int) bool {
		if maxs.Has(data) {
			n.list.Append(data)
		}
		return false
	})
	return n
}

// Diff 差集
func (s *Set) Diff(v *Set) *Set {
	n := New(s.comparator)
	if s.Len() == 0 {
		return n
	} else if v.Len() == 0 {
		return s.Clone()
	}

	s.Each(func(data interface{}, i int) bool {
		if !v.Has(data) {
			n.list.Append(data)
		}
		return false
	})
	return n
}

// Subset 检查v是否是s的子集
func (s *Set) Subset(v *Set) bool {
	if v.Len() == 0 {
		return true
	}
	if v.Len() > s.Len() {
		return false
	}
	rt := true
	v.Each(func(data interface{}, i int) bool {
		if !s.Has(data) {
			rt = false
			return true
		}
		return false
	})
	return true
}

// String 实现fmt.Stringer 接口
func (s *Set) String() string {
	members := make([]string, s.Len())
	s.Each(func(data interface{}, index int) bool {
		members[index] = fmt.Sprint(data)
		return false
	})
	return fmt.Sprintf("Set{%s}", strings.Join(members, ", "))
}

// DefaultMatch 默认的比较函数
func DefaultMatch(a, b interface{}) bool {
	return a == b
}

// New 创建一个新集合
func New(match Comparator, mbs ...interface{}) *Set {
	s := &Set{
		list:       list.New(),
		comparator: match,
	}
	if len(mbs) > 0 {
		for _, v := range mbs {
			s.Insert(v)
		}
	}
	return s
}
