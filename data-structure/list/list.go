// Package list 提供了单链表, 双链表, 循环链表的实现
package list

// 单链表

import "errors"

// 错误代码
var (
	ErrNotFound = errors.New("Node not found")
	ErrRemoved  = errors.New("Node Already remove")
)

// Node 链表节点
type Node struct {
	Data interface{}
	next *Node
	list *List
}

// Remove 从链表中移除
func (n *Node) Remove() {
	n.list.Remove(n)
}

// Append 追加到当前节点后面
func (n *Node) Append(data interface{}) (*Node, error) {
	if n.list == nil {
		return nil, ErrRemoved
	}
	return n.list.Append(data, n)
}

// IsOrphan 检查是否是孤儿节点
func (n *Node) IsOrphan() bool {
	return n.list == nil
}

// BelongsTo 检查是否是指定链表的节点
func (n *Node) BelongsTo(l *List) bool {
	return n.list == l
}

// List 代表单向链表
type List struct {
	head *Node
	tail *Node
	Len  int
}

// Append 向单向链表中添加一个元素
// 如果提供一个节点, 将插入到这个节点厚点
// O(1)
func (l *List) Append(data interface{}, n ...*Node) (*Node, error) {
	nn := &Node{Data: data, list: l}
	if l.head == nil {
		// 列表为空
		l.head = nn
		l.tail = nn
	} else if len(n) == 1 {
		if n[0].list != l {
			return nil, ErrNotFound
		}

		// 插入到el后面
		el := n[0]
		if el == l.tail {
			l.tail = nn
		}
		nn.next = el.next
		el.next = nn
	} else {
		// 插入到尾部
		l.tail.next = nn
		l.tail = nn
	}
	l.Len++
	return nn, nil
}

// Remove 从链表中移除元素
// O(1)
func (l *List) Remove(n *Node) error {
	cur := l.head
	if cur == nil {
		return ErrNotFound
	}

	if n.list != l {
		return ErrNotFound
	}

	// is head
	if cur == n {
		if cur.next != nil {
			l.head = cur.next
		} else {
			l.head = nil
			l.tail = nil
		}
		n.list = nil
		n.next = nil
		l.Len--
		return nil
	}

	// 找到前继
	for cur.next != nil {
		if cur.next == n {
			if cur.next == l.tail {
				l.tail = cur
			}
			cur.next = n.next
			n.list = nil
			n.next = nil
			l.Len--
			return nil
		}
		cur = cur.next
	}

	return ErrNotFound
}

// Each 遍历链表
// O(n)
func (l *List) Each(iteratee func(data interface{}, index int) (stop bool)) {
	if l.Len == 0 {
		return
	}
	cur := l.head
	i := 0
	for cur != nil {
		if ret := iteratee(cur.Data, i); ret {
			break
		}
		cur = cur.next
		i++
	}
}

// IsOrphan 检查节点是否是孤儿, 即不属于任何链表
func (l *List) IsOrphan(n *Node) bool {
	return n.list == nil
}

// Has 检查当前链表是否包含该节点
func (l *List) Has(n *Node) bool {
	return n.list == l
}

// New 创建一个单向链表
func New() *List {
	return &List{}
}
